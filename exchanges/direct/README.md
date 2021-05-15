# Direct

- Routes messages to the queues based on the "routing key" specified in the binding definition.
- In order to sent a message to a queue, routing key of a message & queue need to be exactly same.

## Example

- In Producer, we have created an exchange named `X_direct` with type `direct`.
- Than we bind that exchange to newly created queues named, `Q_direct_error` with routing key `error`, `Q_direct_info` with routing key `info` & `Q_direct_warning` with routing key `warning`.
- Than we published the message to the exchange with specific routing keys. Which than push the message to specific queue to which the exchange is binded with the message routing key.

### Direct Exchange Binding

![Binding](https://github.com/prateeksib/rabbitmq-learning/blob/main/images/fanout-direct.png)

### [Producer](https://github.com/prateeksib/rabbitmq-learning/blob/main/exchanges/direct/producer/producer.go)

![Producer](https://github.com/prateeksib/rabbitmq-learning/blob/main/images/direct-producer.png)

### Queued

![Queued](https://github.com/prateeksib/rabbitmq-learning/blob/main/images/direct-queued-msg.png)

### [Consumer](https://github.com/prateeksib/rabbitmq-learning/blob/main/exchanges/direct/consumer/consumer.go)

![Consumer](https://github.com/prateeksib/rabbitmq-learning/blob/main/images/direct-consumer.png)

### Consumed Messages

![Consumed Messages](https://github.com/prateeksib/rabbitmq-learning/blob/main/images/direct-consumed-msg.png)
