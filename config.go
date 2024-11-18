package dbping

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/jackc/pgx/v5"
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
	Driver   DBDriver `kong:"-"`
}

// Kong hook
// see https://github.com/alecthomas/kong#hooks-beforereset-beforeresolve-beforeapply-afterapply-and-the-bind-option
func (config *Config) AfterApply() error {
	if _, err := mysql.ParseDSN(config.DSN); err == nil {
		config.Driver = DBDriverMySQL
	} else if _, err := pgx.ParseConfig(config.DSN); err == nil {
		config.Driver = DBDriverPostgreSQL
	} else {
		return fmt.Errorf("cannot parse DSN - %s", config.DSN)
	}

	return nil
}
