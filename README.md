goodb
=====

[![Tests](https://github.com/manhtai/goodb/actions/workflows/goodb.yaml/badge.svg)](https://github.com/manhtai/goodb/actions/workflows/goodb.yaml)

> A toy database system, written in Go without third-party dependencies!

## Supported commands

```sql
CREATE TABLE
INSERT
SELECT
DELETE
```

## Supported types

```sql
INT
VARCHAR(1-255)
```

## Get started

```sh
go run main.go
```

## Credits

- DB implementation: [The SimpleDB Database System][0]
- Parser: [Writing An Interpreter In Go][1]


[0]: http://www.cs.bc.edu/~sciore/simpledb/
[1]: https://interpreterbook.com/