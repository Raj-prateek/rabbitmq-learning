# Default

- When a new queue is created on a RabbitMQ system, it is implicitly bound to a system exchange called "default exchange", _with a routing key which is the same as the queue name_
- Default exchange has no name(empty string).
- The Type of default exchange is [`direct`](https://github.com/prateeksib/rabbitmq-learning/tree/main/exchanges/direct)
- When sending a message, if exchange name is left empty, it is handled by the `default exchange`.
- It is automatically created, cannot be deleted and it have a special role.

## Example

- In Producer, we have created an exchange named `X_default` with type `default`.
- Than we crated two new queues named, `Q_default_1`, `Q_default_2`.
- Than we published the message to the exchange with specific routing keys. Which than push the message to specific queue to which the message routing key match the exact name of the queue.

### Direct Exchange Binding

![Binding](https://github.com/prateeksib/rabbitmq-learning/blob/main/images/direct-exchange.png)

### [Producer](https://github.com/prateeksib/rabbitmq-learning/blob/main/exchanges/direct/producer/producer.go)

![Producer](https://github.com/prateeksib/rabbitmq-learning/blob/main/images/direct-producer.png)

### Queued

![Queued](https://github.com/prateeksib/rabbitmq-learning/blob/main/images/direct-queued-msg.png)

### [Consumer](https://github.com/prateeksib/rabbitmq-learning/blob/main/exchanges/direct/consumer/consumer.go)

![Consumer](https://github.com/prateeksib/rabbitmq-learning/blob/main/images/direct-consumer.png)

### Consumed Messages

![Consumed Messages](https://github.com/prateeksib/rabbitmq-learning/blob/main/images/direct-consumed-msg.png)
