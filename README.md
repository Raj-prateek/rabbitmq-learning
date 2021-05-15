# rabbitmq-learning

## Content

- Getting Started
- Introduction to Web Interface
- Messaging Systems & Exchanges
- Types of Exchanges

## Getting Started

### What is RabbitMQ?

- **RabbitMq** is a message broker that recieves messages from producers and routes them to one of the consumers.
- It is written in `erlang` and open source. It is simple and very powerful.
- It's feature can be extended using plugins.

### Installation

To install rabbitmq i have used docker. I have used a docker image `rabbitmq:3-management-alpine`. Configured in [docker-compose.yaml](https://github.com/prateeksib/rabbitmq-learning/blob/main/docker-compose.yaml) file.

### Step to run RabbitMQ

- Run command `docker-compose up -d`. It will automatically install the image and run in background.

### Elements of Messaging System

- [**Message**](https://github.com/prateeksib/rabbitmq-learning#messages) : Data that is need to be processed. It can be any thing such as command, query or information. In rmq it is in `string` format. Every message has two parts: `routing information` and `payload`.
- **Producer/Consumer** : It set a of code which creates and consume message. Can be on different tech stack for a same message.
- [**Message Queues**](https://github.com/prateeksib/rabbitmq-learning#queues) : Message Queues is the list of available message sent by producer. Every queue should consists of unique name.
- [**Broker,Router/Exchange**](https://github.com/prateeksib/rabbitmq-learning#exchanges) : Recieve & deliver messages. Brokers are intermediate elements who transmits message from senders to related reciever. Routers is component of message broker which decide the transferring of message to which queue or queues, based on the configuration. Routing elements are called as exchange.
- **Connection Channels** : Used for communication.
- **Binding** : Define the relationship b/w exchange and queue. Binding definition contains arguments like `routing keys` & `headers` that are used to filter messages that will be sent to a bound queue.

Note: A queue can be bound to 1 or more exchange.

## Introduction to Web Interface

To run web interface open [localhost:15672](http://localhost:15672) on a browser. Login with the `guest` as a username and `guest` as a password.

![Web Interface](https://github.com/prateeksib/rabbitmq-learning/blob/main/images/web-interface.png)

There are four tabs we will talk about them one by one:

- **Overview** : It contains general status information about the RabbitMQ Service. It also provide information about operations statistics, port configuration, export & import settings.
- **Connections** : All the connections made by producer and consumer application are listed here.
- **Channels** : These are the virtual connections created by client applications.
- **Exchanges** : They are the message routing elements of rmq. All exchanges are always created at the time of installation.
- **Queues** : These are the actual list that holds the available messages.
- **Admin** : It is a user access control management section. We can manage users, virtual hosts, feature flags, policies, limits and clusters.

## Messaging Systems & Exchanges

### Messages

#### Attributes of Messages

- **Routing Key** : Single or multiple words that are used when distributing a message to the queue.
- **Headers** : Collection of key value pairs which is used to transfer message with additional information(if needed).
- **Payload** : Actual data that a message carries.
- **Publishing** : Set by a publisher/producer at the time of producing.
- **Expiration** : It is an optional field which consist of expiration time of message to get deleted from a queue automatically. Unit in milliseconds.
- **Delivery Mode** : Persistent or transient, Persistent means data is save on the disk, so data loss doesn't occur is rmq restarts but in transient it can.
- **Priority** : Priority of the message, between 0-255.
- **Message Id** : Optional unique identifier of message set by the publisher.
- **Reply To** : Optional queue or exchange name used in request-response scenarios.

### Queues

#### Attributes of Queues

- **Name** : Unique queue name, max 255 characters UTF-8 string.
- **Durable** : Whether to preseve or delete this queue when rmq restarts.
- **Auto Delete** : Whether to delete this queue if no one is subscribed to it.
- **Exclusive** : Used only by one connection and deleted when the connection is closed.
- **Max Length** : Maximum number of waiting messages in queue. Overflow behaviour can be set as drop the oldest message or reject the new one.
- **Max Priority** : Maximum number of priority from 0-255 that this queue supports. By default this value is 0.
- **Message TTL** : It is an life time for each message added to the queue. If both message and queue have the ttl set the lowest one will be chosen.
- **Dead-letter exchange** : Name of the exchange that expired or dropped messages will be automatically sent.
- **Binding Configuration** : Association between queues and exchanges. A queue must be bound to a exchange, in order to recieve message from it.

### Exchanges

Exchanges are the messages router elements of RabbitMQ

#### Attributes of Exchanges

- **Name** : Unique name.
- **Types** : `default`, `fanout`, `direct`, `topic` or `headers`.
- **Durable** : Same as queues durability. Durable exchanges survive after a service restarts.
- **Auto Delete** : Whether to delete this exchange if no bound queue is left.
- **Alternate Exchange** : Unroutable messages will be sent using alternate exchange.
