FROM bitnami/kafka:latest

WORKDIR /

COPY config/ /config/
COPY bin/setup.sh .

EXPOSE 9092

CMD ["./setup.sh"]
