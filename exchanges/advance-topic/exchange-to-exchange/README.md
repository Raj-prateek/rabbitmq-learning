# Exchange bind to exchange

- Like binding a queue to an exchange, in rabbit mq we can bind exchange to another exchange.
- Binding and messaging routing rules are same.
- When an exchange is bound to another exchange, messages from the source exchange are routed to the destination exchange using the binding configuration.
- Finally, destination exchange routes that message to its bind queue.

## Example

- We have created an exchange named `X_one` with type `direct` & `X_two` with type `direct`.
- Than we bind that exchange to newly created queues named, `Q_one` with routing key `one`, `Q_two` with routing key `two`.
- Now we have binded exchange `X_one` to `X_two` on key `two`.
- Than we published the message to the exchange `X_one` with key `two`. Which than route message to exchange to `X_two` and than routed to queue `Q_one`.

### Exchange to Exchange Binding

![Binding](https://github.com/prateeksib/rabbitmq-learning/blob/main/images/exchange-to-exchange.png)

### [Producer](https://github.com/prateeksib/rabbitmq-learning/blob/main/exchanges/advance-topic/exchange-to-exchange/producer/producer.go)

![Producer](https://github.com/prateeksib/rabbitmq-learning/blob/main/images/exchange-to-exchange-producer.png)

### Queued

![Queued](https://github.com/prateeksib/rabbitmq-learning/blob/main/images/exchange-to-exchange-queued-msg.png)

### [Consumer](https://github.com/prateeksib/rabbitmq-learning/blob/main/exchanges/advance-topic/exchange-to-exchange/consumer/consumer.go)

![Consumer](https://github.com/prateeksib/rabbitmq-learning/blob/main/images/exchange-to-exchange-consumer.png)

### Consumed Messages

![Consumed Messages](https://github.com/prateeksib/rabbitmq-learning/blob/main/images/exchange-to-exchange-consumed-msg.png)
