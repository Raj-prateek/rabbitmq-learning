# Fanout

- It is a simple type of exchange, in which all the incoming messages to the exchange is passed to all the queues bind to that exchange.
- It simply ignore the routing information and doesn't perform any filtering.

## Example

- In Producer itself, we created an exchange named `X_fanout` with type `fanout`.
- Than we bind that exchange to newly created queues named, `Q_fanout_1` & `Q_fanout_2`.
- Than we published the message to the exchange. Which than push the message to both the queues.

### Fanout Exchange Binding

![Binding](https://github.com/prateeksib/rabbitmq-learning/blob/main/images/fanout-exchange.png)

### [Producer](https://github.com/prateeksib/rabbitmq-learning/blob/main/exchanges/fanout/producer/producer.go)

![Producer](https://github.com/prateeksib/rabbitmq-learning/blob/main/images/fanout-producer.png)

### Queued

![Queued](https://github.com/prateeksib/rabbitmq-learning/blob/main/images/fanout-queued-msg.png)

### [Consumer](https://github.com/prateeksib/rabbitmq-learning/blob/main/exchanges/fanout/consumer/consumer.go)

![Consumer](https://github.com/prateeksib/rabbitmq-learning/blob/main/images/fanout-queued-msg.png)

### Consumed Messages

![Consumed Messages](https://github.com/prateeksib/rabbitmq-learning/blob/main/images/fanout-consumed-msg.png)
