extern crate dotenv;

use dotenv::dotenv;
use kafka::client::{FetchOffset, GroupOffsetStorage};
use kafka::consumer::Consumer;
use kafka::error::Error as KafkaError;
use std::env;
use std::string::String;
use redis;

fn redis_connect() -> redis::Connection {
    let redis_host = match env::var("REDIS_HOST") {
        Ok(val) => val,
        Err(_) => format!("{}", "localhost"),
    };

    let redis_port = match env::var("REDIS_PORT") {
        Ok(val) => val,
        Err(_) => format!("{}", "localhost"),
    };
    
    let redis_tls = match env::var("IS_TLS") {
        Ok(_) => "rediss",
        Err(_) => "redis",
    };

    let redis_url = format!("{:}://{:}:{:}/", redis_tls, redis_host, redis_port);

    redis::Client::open(redis_url)
        .expect("invalid URL")
        .get_connection()
        .expect("unable to connect")
}

fn start_consumer(topic: String, brokers: Vec<String>) -> Result<(), KafkaError> {
    let mut conn_redis = redis_connect();

    let mut conn_kafka = Consumer::from_hosts(brokers)
        .with_topic(topic)
        .with_fallback_offset(FetchOffset::Earliest)
        .with_offset_storage(GroupOffsetStorage::Kafka)
        .create()?;

    loop {
        let mss = conn_kafka.poll()?;

        if mss.is_empty() {
            println!("No messages")
        }

        for ms in mss.iter(){
            for m in ms.messages(){

                let key = String::from_utf8_lossy(&m.key);
                let value = String::from_utf8_lossy(&m.value);

                let _: () = redis::cmd("SET")
                    .arg(key.to_string())
                    .arg(value.to_string())
                    .query(&mut conn_redis)
                    .expect("Not able to set key");
            }
            let _ = conn_kafka.consume_messageset(ms);
        }
        conn_kafka.commit_consumed()?;
    }
}

//KAFKA_URL=localhost
//KAFKA_PORT=9092
//KAFKA_TOPIC=commandss


fn main() {
    dotenv().ok();


    let kafka_url = match env::var("KAFKA_URL") {
        Ok(val) => val,
        Err(_) => format!("{}", "localhost"),
    };

    let kafka_port = match env::var("KAFKA_PORT") {
        Ok(val) => val,
        Err(_) => format!("{}", "localhost"),
    };
    
    let kafka_topic = match env::var("KAFKA_TOPIC") {
        Ok(val) => val,
        Err(_) => format!("{}", "localhost"),
    };

    let kafka_conn_url = format!("{:}:{:}", kafka_url, kafka_port);

    if let Err(e) = start_consumer(kafka_topic, vec![kafka_conn_url]) {
        println!("{:}", e)
    }
}
