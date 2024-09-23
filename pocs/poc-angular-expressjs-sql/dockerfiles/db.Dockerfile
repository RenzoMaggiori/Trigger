FROM mysql:8

COPY ./backend/src/db/init.sql /docker-entrypoint-initdb.d/

EXPOSE 3306