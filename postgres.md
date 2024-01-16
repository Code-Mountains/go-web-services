## Commands:
```
CREATE DATABASE readinglist;

CREATE ROLE readinglist WITH LOGIN PASSWORD 'pa55w0rd';

\c readinglist;

CREATE TABLE IF NOT EXISTS books (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    title text NOT NULL,
    published integer NOT NULL,
    pages integer NOT NULL,
    genres text[] NOT NULL,
    rating real NOT NULL,
    version integer NOT NULL DEFAULT 1
);

GRANT SELECT, INSERT, UPDATE, DELETE on books TO readinglist;

GRANT USAGE, SELECT on SEQUENCE books_id_seq TO readinglist;
```

## With Output:
```
postgres=# CREATE DATABASE readinglist;
CREATE DATABASE

postgres=# CREATE ROLE readinglist WITH LOGIN PASSWORD 'pa55w0rd';
CREATE ROLE

postgres=# \c readinglist;
psql (14.10 (Ubuntu 14.10-0ubuntu0.22.04.1), server 16.1 (Debian 16.1-1.pgdg120+1))
WARNING: psql major version 14, server major version 16.
         Some psql features might not work.
You are now connected to database "readinglist" as user "postgres".

readinglist=# CREATE TABLE IF NOT EXISTS books (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    title text NOT NULL,
    published integer NOT NULL,
    pages integer NOT NULL,
    genres text[] NOT NULL,
    rating real NOT NULL,
    version integer NOT NULL DEFAULT 1
);
CREATE TABLE

readinglist=# GRANT SELECT, INSERT, UPDATE, DELETE on books TO readinglist;
GRANT

readinglist=# GRANT USAGE, SELECT on SEQUENCE books_id_seq TO readinglist;
GRANT
```

## Connection String / DSN for Go app to connect to Postgres
```
export READINGLIST_DB_DSN='postgres://readinglist:pa55w0rd@localhost/readinglist?sslmode=disable'
```


## Export PSQL password to .bashrc for psql cli access without pwd prompt
```
export PGPASSWORD='mysecretpassword' >> .bashrc
```