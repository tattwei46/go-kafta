package main

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"git.davidcheah.com/go-kafka/producer"
)

func main() {
	fmt.Println("listening...")
	if err := subscribeKafka(); err != nil {
		panic(err.Error())
	}
}

func subscribeKafka() error {
	configMap := kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "group-id-1",
		"auto.offset.reset": "earliest",
	}

	c, err := kafka.NewConsumer(&configMap)
	if err != nil {
		return err
	}

	topics := []string{producer.JobsTopic}
	if err := c.SubscribeTopics(topics, nil); err != nil {
		return err
	}

	for {
		msg, err := c.ReadMessage(-1)
		if err != nil {
			break
		}
		fmt.Printf("Received from kafka: %s, %s\n", msg.TopicPartition, string(msg.Value))

		job, err := getJob(msg.Value)
		if err != nil {
			fmt.Println("an error occurs when getting job, ", err.Error())
		} else {
			fmt.Println("Received job: ",job)
		}
	}

	return nil

}

func getJob(b []byte ) (producer.Job,error){
	var job producer.Job
	err := json.Unmarshal(b, &job)
	return job, err
}
