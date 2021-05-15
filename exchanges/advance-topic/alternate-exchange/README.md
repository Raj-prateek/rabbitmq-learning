# Alternate exchange

- Some of the message published to the exchange may be not suitable to route to any of the bound queues.
- These are unrouted messages.
- They are discarded by the exchange, so they are lost.
- In order to collect these messages, an `alternative exchange` can be defined for any exchange.
- Alternative exchange can be define by setting the `alternate-exchange` key for an exchange.
- Any unrouted message is finally sent to defined `alternate-exchange`.
- Any existing exchange can be set as an `alternate exchange`for another exchange.
- [Fanout exchange](https://github.com/prateeksib/rabbitmq-learning/tree/main/exchanges/fanout) which doesn't used any filter, are good for using as an `alternate exchange`.

## Example

- We have created an exchange named `X_alternate_one` with type `direct` with alternative exchange as `X_alternate_two` & `X_alternate_two` with type `fanout`.
- Than we bind that exchange to newly created queues named, `Q_alternate_one` & `Q_alternate_two` with routing key `one` & `two` respectively, `Q_alternate_unrouting` with exchange `X_alternate_two(fanout)`.
- Than we published the message to the exchange `X_alternate_one` with routing key `one`, `two` & `three`. From which message with routing key `one` & `two` is routed to `Q_alternate_one` & `Q_alternate_two` respectively, on the other hand message with routing key `three` is routed to the exchange `X_alternate_two` which than routed to the queue `Q_alternate_unrouting`.

### Alternate Exchange

![Binding](https://github.com/prateeksib/rabbitmq-learning/blob/main/images/alternate-exchange.png)

### [Producer](https://github.com/prateeksib/rabbitmq-learning/blob/main/exchanges/advance-topic/alternate-exchange/producer/producer.go)

![Producer](https://github.com/prateeksib/rabbitmq-learning/blob/main/images/alternate-exchange-producer.png)

### Queued

![Queued](https://github.com/prateeksib/rabbitmq-learning/blob/main/images/alternate-exchange-queued-msg.png)

### [Consumer](https://github.com/prateeksib/rabbitmq-learning/blob/main/exchanges/advance-topic/alternate-exchange/consumer/consumer.go)

![Consumer](https://github.com/prateeksib/rabbitmq-learning/blob/main/images/alternate-exchange-consumer.png)

### Consumed Messages

![Consumed Messages](https://github.com/prateeksib/rabbitmq-learning/blob/main/images/alternate-exchange-consumed-msg.png)
