# rabbitmq-learning

## Content

- [Getting Started](https://github.com/prateeksib/rabbitmq-learning#getting-started)
- [Introduction to Web Interface](https://github.com/prateeksib/rabbitmq-learning#introduction-to-web-interface)
- [Messaging Systems & Exchanges](https://github.com/prateeksib/rabbitmq-learning#messaging-systems--exchanges)
- [Types of Exchanges](https://github.com/prateeksib/rabbitmq-learning/tree/main/exchanges)
- [Push vs Pull](https://github.com/prateeksib/rabbitmq-learning#push-vs-pull)
- [Work/Task Queues(Competing Consumers Pattern)](https://github.com/prateeksib/rabbitmq-learning#work-task-queues-competing-consumers-pattern)
- [Publish/Subscribe Pattern](https://github.com/prateeksib/rabbitmq-learning#publish-subscribe-pattern)
- [Request-Reply Pattern](https://github.com/prateeksib/rabbitmq-learning#request-reply-pattern)

## Getting Started

### What is RabbitMQ?

- **RabbitMq** is a message broker that recieves messages from producers and routes them to one of the consumers.
- It is written in `erlang` and open source. It is simple and very powerful.
- It's feature can be extended using plugins.

### Installation

To install rabbitmq we will used a docker image `rabbitmq:3-management-alpine`. Configured in [docker-compose.yaml](https://github.com/prateeksib/rabbitmq-learning/blob/main/docker-compose.yaml) file.

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

### [Exchanges](https://github.com/prateeksib/rabbitmq-learning/tree/main/exchanges)

Exchanges are the messages router elements of RabbitMQ

#### Attributes of Exchanges

- **Name** : Unique name.
- **Types** : `default`, `fanout`, `direct`, `topic` or `headers`.
- **Durable** : Same as queues durability. Durable exchanges survive after a service restarts.
- **Auto Delete** : Whether to delete this exchange if no bound queue is left.
- **Alternate Exchange** : Unroutable messages will be sent using alternate exchange.

## Push vs Pull

There are mainly two approach to consume message from a queue.

- ### Push

  - Consumer application subscribes to the queue and waits for messages.
  - If there is already a message on the queue.
  - Or when the new message arrives, it automatically sent _(pushed)_ to the consumer application.
  - This is the suggested way of getting messages from a queue.

- ### Pull

  - Consumer application doesn't subscribe to queue.
  - But it constantly checks(polls) the queue for new messages.
  - If there is a message available on the queue, it is manually fetched (pulled) by the consumer application.
  - Even though the pull mode is not recommended, it is the only solution when there is no live connection between message broker and consumer applications.

## Work/Task Queues(Competing Consumers Pattern)

- Work queues are used to distribute tasks among multiple workers.
- Producers add tasks to a queue and these tasks are distributed to multiple worker applications.
- Pull or push models can be used to distribute tasks among the workers.
  - **Pull Model**: Workers get a message from the queue when they are available to perform a task.
  - **Push Model**: Message broker system sends _(pushes)_ messages to the available workers automatically.

### Issues

- Messages must not be lost if the broker crashes.
- Every message must be delivered and processed for exactly one time.
- If the worker application crashes or fails to process a task, it must be re-added to the queue and delivered later.
- There must be a limit for retrying failed tasks, otherwise the queue may be filled up by the constantly failing tasks.
- Task workload should be fairly distributed among the workers.

![Worker](https://github.com/prateeksib/rabbitmq-learning/blob/main/images/worker-simultaneously.png)

![Worker with timestamp](https://github.com/prateeksib/rabbitmq-learning/blob/main/images/worker-simultaneously-with-time.png)

## Publish/Subscribe Pattern

- Publish-subscribe pattern is used to deliver the same message to all the subscriber.
- There may be one or more subscribers.
- In publish-subscriber pattern, there is one queue for each subscriber.
- Message is delivered to each of these queues.
- So, each subscriber gets a copy of the same messages.
- This pattern is mostly used for publishing event notifications.

## Request-Reply Pattern

- Request-reply pattern is used when the publisher of the message, which is called requestor, needs to get the response for its message.
- The request message mostly contains a query or a command.
- Using request-reply pattern, remote procedure call _(RPC)_ scenarios can also be implemented.
- In request-reply patterns, there are at least two queues:
  - 1 for the requests.
  - 1 for the replies. This queue also named as the callback queue for _RPC_ scenarios.

## Priority Queues - Message Priorities

- Configuring RabbitMQ queues and channels for message priorities:
  - Set the max priority value of the queue by setting the `x-max-priority` argument. It can be 0 to 255. 0 means queue doesn't supports priorities. Values between 1 & 10 are recommended.
  - Set the priority of the message using the `basicProperties.Priority`.
  - Messages must stay in the queue for a while, to be ordered properly.
  - Use `channel.Qos(...)` to set the `Prefetch Count` to 1.
    _This will configure a worker's channel, not to send a new message until the worker finishes the last message_
  - Set `Auto ACK` value to `false` while subscribing to a queue.
  - if you don't configure the channel for not sending a new message until the worker acknowledges the last one, every incoming message will be immediately sent to a worker.
