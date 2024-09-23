FROM postgres:15

COPY ./backend/db/init.sql /docker-entrypoint-initdb.d/

EXPOSE 5432
