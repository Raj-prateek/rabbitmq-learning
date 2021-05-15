# Headers

- Uses message header in order to route messages bounded to a queues.
- Ignore the routing key value of the message.
- A message can have many different header and with many different values.
- While binding to this type of exchange, every queue specifies which headers a message may contains and whether it requires "all" or "any" of them to be exist in the queue.
- `x-match` is the special header key whose value can be "all" or "any". It determines the "match all" or "match any" logic for the matching process.

## Example

- We have created an exchange named `X_headers` with type `headers`.
- Than we bind that exchange to newly created queues named, `Q_headers_1` with headers `{"x-match": "all", "job": "convert", "format": "jpeg"}` & `Q_headers_2` with headers `{"x-match": "any", "job": "convert", "format": "jpeg"}`.
- Than we published the message to the exchange. Which than push the message to specific queue to which the header exact match or any of the key match depending on the header key binded with the queue.

### headers Exchange Binding

![Binding](https://github.com/prateeksib/rabbitmq-learning/blob/main/images/headers-exchange.png)

### [Producer](https://github.com/prateeksib/rabbitmq-learning/blob/main/exchanges/headers/producer/producer.go)

![Producer](https://github.com/prateeksib/rabbitmq-learning/blob/main/images/headers-producer.png)

### Queued

![Queued](https://github.com/prateeksib/rabbitmq-learning/blob/main/images/headers-queued-msg.png)

### [Consumer](https://github.com/prateeksib/rabbitmq-learning/blob/main/exchanges/headers/consumer/consumer.go)

![Consumer](https://github.com/prateeksib/rabbitmq-learning/blob/main/images/headers-consumer.png)

### Consumed Messages

![Consumed Messages](https://github.com/prateeksib/rabbitmq-learning/blob/main/images/headers-consumed-msg.png)
