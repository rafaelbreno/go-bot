# go-bot
- Simple twitch bot made with Go - Study purpose
- I'm following the official [Twitch Chatbot Docs](https://dev.twitch.tv/docs/irc/guide)
- [AuthToken](https://twitchapps.com/tmi/)
- [Application Architecture](https://whimsical.com/gobot-UhQLa6aXBkAXd4tSmJn5EZ)

### Deploy
- Docker

### TODOv1:
- [x] Stablish a Bot
  - [x] IRC Connection
  - [X] Commands
  - [x] Database (JSON)

### TODOv2:
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

### TODOv3:
- [x] Study the difference between:
  - [x] RabbitMQ and Kafka
  - [x] HTTP and gRPC
- [x] Up Kafka using Docker
- [x] API
  - [x] Implement the communication between API and Kafka/RabbitMQ
  - [x] Produce - Kafka

### TODOv4:
- [x] Queue Manager
  - [x] Create microservice
  - [x] Implement the communication between API and Kafka/RabbitMQ
  - [x] Implement Redis
  - [x] Consume - Kafka

### TODOv5:
- [ ] Auth service
  - [ ] Redis Connection
  - [ ] Postgres Connection
  - [ ] Endpoints
    - [ ] SignUP
    - [ ] SignIn
    - [ ] SignOut
    - [ ] CheckAuth
  - [ ] 

### TODO(v?):
- Queue Mgr
  - [ ] Insert/Update/Delete Redis rows
  - [ ] Implement HealthCheck
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
