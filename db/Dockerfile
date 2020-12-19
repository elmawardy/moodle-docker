FROM postgres:12.5
ENV POSTGRES_DB moodle
ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD 123456

COPY ./database.sql /tmp/
COPY restore-database.sh /docker-entrypoint-initdb.d/restore-database.sh
RUN sed -i 's/\r$//g' /docker-entrypoint-initdb.d/restore-database.sh
RUN chmod 777 /docker-entrypoint-initdb.d/restore-database.sh

EXPOSE 5432