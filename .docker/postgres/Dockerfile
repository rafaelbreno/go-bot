FROM postgres:13-alpine

COPY /.docker/postgres/conf/pg_hba.conf /var/lib/postgresql/data/pg_hba.conf

CMD ["postgres"]

EXPOSE 5432
