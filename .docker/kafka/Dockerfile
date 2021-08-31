FROM bitnami/kafka:latest

WORKDIR /

COPY config/server.properties /config/server.properties
COPY bin/setup.sh .

EXPOSE 9092

CMD ["./setup.sh"]
