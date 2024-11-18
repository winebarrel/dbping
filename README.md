# dbping

[![CI](https://github.com/winebarrel/dbping/actions/workflows/ci.yml/badge.svg)](https://github.com/winebarrel/dbping/actions/workflows/ci.yml)

PING for DB.

## Usage

```
sage: dbping <dsn> [flags]

Arguments:
  <dsn>    DSN to connect to.
             - MySQL: https://pkg.go.dev/github.com/go-sql-driver/mysql#readme-dsn-data-source-name
             - PostgreSQL: https://pkg.go.dev/github.com/jackc/pgx/v5/stdlib#pkg-overview

Flags:
  -h, --help       Show help.
      --version
```

```
$ dbping 'root@tcp(127.0.0.1:13306)/mysql'
OK 4.152041ms
OK 1.601542ms
OK 1.205208ms
[ERROR] driver: bad connection
[ERROR] dial tcp 127.0.0.1:13306: connect: connection refused
[ERROR] dial tcp 127.0.0.1:13306: connect: connection refused
OK 6.901792ms
OK 2.089375ms
OK 2.137875ms
...
```
