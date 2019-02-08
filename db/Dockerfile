FROM mysql

ENV MYSQL_DATABASE=codecamp \
    MYSQL_ROOT_PASSWORD=codecamp

ADD schema.sql /docker-entrypoint-initdb.d

EXPOSE 3306
