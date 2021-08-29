# go-bot
- Simple twitch bot made with Go - Study purpose
- I'm following the official [Twitch Chatbot Docs](https://dev.twitch.tv/docs/irc/guide)
- [AuthToken](https://twitchapps.com/tmi/)
- [Application Architecture](https://whimsical.com/gobot-UhQLa6aXBkAXd4tSmJn5EZ)

### Deploy
- Docker

### TODO(v1):
- [x] Stablish a Bot
  - [x] IRC Connection
  - [X] Commands
  - [x] Database (JSON)

### TODO(v2)
- [x] Develop API
  - [x] Establish API base
    - [x] Context
    - [x] Logger (uber-zap)
    - [x] WebFramework (fiber)
    - [x] ORM/DB Connection (GORM)
  - [x] Define DB
    - [x] SQL - PostgreSQL
      - [x] Write Migrations
      - [x] Define Entities

### TODO(v3):
- [ ] Study the difference between:
  - [ ] RabbitMQ and Kakfa
  - [ ] HTTP and gRPC
- [ ] Up Kafka/RabbitMQ using Docker
- [ ] API
  - [ ] Implement the communication between API and Kafka/RabbitMQ
  - [ ] Produce - RabbitMQ/Kafka
- [ ] Queue Manager
  - [ ] Create microservice
  - [ ] Implement the communication between API and Kafka/RabbitMQ
  - [ ] Implement Redis
  - [ ] Consume - RabbitMQ/Kafka
  - [ ] Insert/Update/Delete Redis rows
  - [ ] Implement HealthCheck

### TODO(v?):
- [ ] Local deploy Docker
- [ ] Front-end
  - [ ] Choose one
    - [ ] React (with Rescript)
    - [ ] Vue
    - [ ] Svelte <- more likely
  - [ ] OAuth
- [ ] Prometheus
- [ ] noSQL - Redis
  - [ ] Define an Entity
  - [x] Define key layout
    - [stream][user]
- [ ] Implement HealthCheck
