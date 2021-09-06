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
    let mut con = Consumer::from_hosts(brokers)
        .with_topic(topic)
        .with_fallback_offset(FetchOffset::Earliest)
        .with_offset_storage(GroupOffsetStorage::Kafka)
        .create()?;

    loop {
        let mss = con.poll()?;

        if mss.is_empty() {
            println!("No messages")
        }

        for ms in mss.iter(){
            for m in ms.messages(){
                println!("{}:{}@{}: {:?} - {:?}", ms.topic(), ms.partition(), m.offset, m.key, m.value);
            }

        }
    }
}

//KAFKA_URL=localhost
//KAFKA_PORT=9092
//KAFKA_TOPIC=commandss


fn main() {
    dotenv().ok();

    let mut conn = redis_connect();

    let _: () = redis::cmd("SET")
        .arg("key")
        .arg("value")
        .query(&mut conn)
        .expect("'key' not found");

    let found: String = redis::cmd("GET")
        .arg("key")
        .query(&mut conn)
        .expect("'key' not found");

    println!("{}", found);

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
