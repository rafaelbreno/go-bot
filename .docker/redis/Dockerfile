FROM redis:6-buster

COPY /.docker/redis/conf/redis.conf /usr/local/etc/redis/redis.conf

CMD ["redis-server", "/usr/local/etc/redis/redis.conf"]

EXPOSE 6379