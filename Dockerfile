FROM postgres:alpine
COPY ./db-structure/script.sql /docker-entrypoint-initdb.d/init.sql
EXPOSE 5432
