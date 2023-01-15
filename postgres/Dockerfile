FROM postgres:15.1

COPY schema.sql /docker-entrypoint-initdb.d/000-schema.sql
COPY functions.sql /docker-entrypoint-initdb.d/001-function.sql