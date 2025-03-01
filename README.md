# dbping

[![CI](https://github.com/winebarrel/dbping/actions/workflows/ci.yml/badge.svg)](https://github.com/winebarrel/dbping/actions/workflows/ci.yml)

PING for DB.

## Usage

```
Usage: dbping <dsn> [flags]

Arguments:
  <dsn>    DSN to connect to.
             - MySQL: https://pkg.go.dev/github.com/go-sql-driver/mysql#readme-dsn-data-source-name
             - PostgreSQL: https://pkg.go.dev/github.com/jackc/pgx/v5/stdlib#pkg-overview

Flags:
  -h, --help            Show help.
  -i, --interval=3      Interval seconds.
  -q, --query=STRING    Query to run.
      --iam-auth        Use IAM authentication.
      --version
```

```
$ dbping 'root@tcp(127.0.0.1:13306)/mysql'
PING 4ms
PING 1ms
PING 1ms
[ERROR] driver: bad connection
[ERROR] dial tcp 127.0.0.1:13306: connect: connection refused
[ERROR] dial tcp 127.0.0.1:13306: connect: connection refused
PING 6ms
PING 2ms
PING 2ms
...

% dbping 'root@tcp(127.0.0.1:13306)/mysql' -q 'select connection_id()'
11 3ms
11 2ms
11 2ms
...

% dbping 'postgres://postgres@localhost:15432' -q 'select pg_backend_pid()'
68 13ms
68 3ms
68 2ms
...
```
