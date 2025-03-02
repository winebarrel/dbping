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
  -t, --timeout=3       Timeout seconds.
  -q, --query=STRING    Query to run.
      --iam-auth        Use IAM authentication.
      --version
```

```
$ dbping 'root@tcp(127.0.0.1:13306)/mysql'
12:27:43 | PING 4ms
12:27:47 | PING 1ms
12:27:50 | PING 1ms
[ERROR] driver: bad connection
[ERROR] dial tcp 127.0.0.1:13306: connect: connection refused
[ERROR] dial tcp 127.0.0.1:13306: connect: connection refused
12:28:02 | PING 6ms
12:28:05 | PING 2ms
12:28:08 | PING 2ms
...

% dbping 'root@tcp(127.0.0.1:13306)/mysql' -q 'select connection_id()'
12:29:06 | 11 3ms
12:29:09 | 11 2ms
12:29:12 | 11 2ms
...

% dbping 'postgres://postgres@localhost:15432' -q 'select pg_backend_pid()'
12:29:39 | 68 13ms
12:29:42 | 68 3ms
12:29:45 | 68 2ms
...
```
