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

## Running Kafka Cluster in docker
### Build docker images of Zookeeper and Kafka
```
docker-compose build
```
### Run docker images of Zookeeper and Kafka
```
MY_IP=192.168.1.106 docker-compose up  
```
## Running Kafka Cluster Manually
### Start Zookeeper
```cassandraql
bin/zookeeper-server-start.sh config/zookeeper.properties
```

### Start Multiple Kafka
```cassandraql
bin/kafka-server-start.sh config/server.properties
```
### Create Topic
```cassandraql
 bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 3 --partitions 1 --topic replicationtopic
```

### Create Producer
```cassandraql
bin/kafka-console-producer.sh --broker-list localhost:9092, localhost:9093, localhost:9094 --topic replicationtopic
```

### Create Consumer
```cassandraql
bin/kafka-console-consumer.sh --bootstrap-server :9092 --topic replicationtopic
```