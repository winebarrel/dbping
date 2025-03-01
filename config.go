package dbping

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type DBDriver string

const (
	DBDriverMySQL      DBDriver = "mysql"
	DBDriverPostgreSQL DBDriver = "pgx"
)

type Config struct {
	DSN      string   `kong:"arg='',required,help='DSN to connect to. \n - MySQL: https://pkg.go.dev/github.com/go-sql-driver/mysql#readme-dsn-data-source-name \n - PostgreSQL: https://pkg.go.dev/github.com/jackc/pgx/v5/stdlib#pkg-overview'"`
	Interval uint     `kong:"short='i',default='3',help='Interval seconds.'"`
	Query    string   `kong:"short='q',help='Query to run.'"`
	IAMAuth  bool     `kong:"help='Use IAM authentication.'"`
	Driver   DBDriver `kong:"-"`
}

// Kong hook
// see https://github.com/alecthomas/kong#hooks-beforereset-beforeresolve-beforeapply-afterapply-and-the-bind-option
func (cfg *Config) AfterApply() error {
	if _, err := mysql.ParseDSN(cfg.DSN); err == nil {
		cfg.Driver = DBDriverMySQL
	} else if _, err := pgx.ParseConfig(cfg.DSN); err == nil {
		cfg.Driver = DBDriverPostgreSQL
	} else {
		return fmt.Errorf("cannot parse DSN - %s", cfg.DSN)
	}

	return nil
}

func (cfg *Config) OpenDB() (*sql.DB, error) {
	var connector driver.Connector
	var err error

	switch cfg.Driver {
	case DBDriverMySQL:
		connector, err = cfg.getMySQLConnector()
	case DBDriverPostgreSQL:
		connector, err = cfg.getPostgreSQLConnector()
	default:
		err = fmt.Errorf("unimplemented driver - %s", cfg.Driver)
	}

	if err != nil {
		return nil, err
	}

	db := sql.OpenDB(connector)
	db.SetConnMaxLifetime(0)
	db.SetConnMaxIdleTime(0)
	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(1)

	return db, nil
}

func (cfg *Config) getMySQLConnector() (driver.Connector, error) {
	mycfg, err := mysql.ParseDSN(cfg.DSN)

	if err != nil {
		return nil, err
	}

	if cfg.IAMAuth {
		hostPort := strings.SplitN(mycfg.Addr, ":", 2)
		host, err := resolveCNAME(hostPort[0])

		if err != nil {
			return nil, err
		}

		port := hostPort[1]
		endpoint := host + ":" + port
		user := mycfg.User

		bc := func(ctx context.Context, mc *mysql.Config) error {
			token, err := buildIAMAuthToken(ctx, endpoint, user)

			if err != nil {
				return err
			}

			mc.Passwd = token
			return nil
		}

		err = mycfg.Apply(mysql.BeforeConnect(bc))

		if err != nil {
			return nil, err
		}

		mycfg.AllowCleartextPasswords = true

		if mycfg.TLSConfig == "" {
			mycfg.TLSConfig = "preferred"
		}
	}

	return mysql.NewConnector(mycfg)
}

func (cfg *Config) getPostgreSQLConnector() (driver.Connector, error) {
	opts := []stdlib.OptionOpenDB{}
	pgcfg, err := pgx.ParseConfig(cfg.DSN)

	if err != nil {
		return nil, err
	}

	if cfg.IAMAuth {
		host, err := resolveCNAME(pgcfg.Config.Host)

		if err != nil {
			return nil, err
		}

		endpoint := fmt.Sprintf("%s:%d", host, pgcfg.Config.Port)
		user := pgcfg.Config.User

		opts = append(opts, stdlib.OptionBeforeConnect(func(ctx context.Context, cc *pgx.ConnConfig) error {
			token, err := buildIAMAuthToken(ctx, endpoint, user)

			if err != nil {
				return err
			}

			cc.Password = token
			return nil
		}))
	}

	connector := stdlib.GetConnector(*pgcfg, opts...)
	return connector, nil
}
