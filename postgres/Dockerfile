FROM postgres:15.1

RUN localedef -i fr_FR -c -f UTF-8 -A /usr/share/locale/locale.alias fr_FR.UTF-8

COPY schema.sql /docker-entrypoint-initdb.d/000-schema.sql
COPY functions.sql /docker-entrypoint-initdb.d/001-function.sql