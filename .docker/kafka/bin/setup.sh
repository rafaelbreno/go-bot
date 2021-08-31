kafka-server-start.sh /config/server.properties

kafka-topics.sh --create --topic quickstart-events --bootstrap-server 127.0.0.1:9000
