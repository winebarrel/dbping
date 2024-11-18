package dbping

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/jackc/pgx/v5"
)

type DBDriver string

const (
	DBDriverMySQL      DBDriver = "mysql"
	DBDriverPostgreSQL DBDriver = "pgx"
)

type DBConfig struct {
	DSN    string   `kong:"arg='',required,help='DSN to connect to. \n - MySQL: https://pkg.go.dev/github.com/go-sql-driver/mysql#readme-dsn-data-source-name \n - PostgreSQL: https://pkg.go.dev/github.com/jackc/pgx/v5/stdlib#pkg-overview'"`
	Driver DBDriver `kong:"-"`
}

// Kong hook
// see https://github.com/alecthomas/kong#hooks-beforereset-beforeresolve-beforeapply-afterapply-and-the-bind-option
func (config *DBConfig) AfterApply() error {
	if _, err := mysql.ParseDSN(config.DSN); err == nil {
		config.Driver = DBDriverMySQL
	} else if _, err := pgx.ParseConfig(config.DSN); err == nil {
		config.Driver = DBDriverPostgreSQL
	} else {
		return fmt.Errorf("cannot parse DSN - %s", config.DSN)
	}

	return nil
}

func (config *DBConfig) Open() (*sql.DB, error) {
	db, err := sql.Open(string(config.Driver), config.DSN)

	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(0)
	db.SetConnMaxIdleTime(0)
	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(1)

	return db, nil
}
