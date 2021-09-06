extern crate dotenv;

use dotenv::dotenv;
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
}
