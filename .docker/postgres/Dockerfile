FROM postgres:alpine

COPY /.docker/postgres/config/pg_hba.conf /var/lib/postgresql/data/pg_hba.conf

CMD ["postgres"]

EXPOSE 5432