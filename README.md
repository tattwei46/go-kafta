# Golang With Kafka Example

## Download Kafka
https://kafka.apache.org/downloads

## Start Zookeeper
```
bin/zookeeper-server-start.sh config/zookeeper.properties
```

## Start Kafka
```
bin/kafka-server-start.sh config/server.properties
```

## Start Producer
```
go build /golang-series/go-kafka/producer/main.go
```

## Start Consumer
```
go build /golang-series/go-kafka/consumer/main.go
```

## Send Message
```
curl --location --request POST 'http://0.0.0.0:8080/go-kafka/v1/jobs' \
--header 'Content-Type: application/json' \
--data-raw '{
	"title": "test",
	"description": "description"
}'
```
