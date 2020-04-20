package main

import (
	"encoding/json"
	"fmt"
	"git.davidcheah.com/go-kafka/producer"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()

	base := router.Group("/go-kafka/v1")

	createRoute(base)

	host := fmt.Sprintf("%s:%d", "0.0.0.0", 8080)

	if err := router.Run(host); err != nil {
		fmt.Printf("\n Error occurs while starting the server")
	}
}

func createRoute(r *gin.RouterGroup) {
	r.POST("/jobs", getJob)
}

func getJob(c *gin.Context) {
	var job producer.Job
	err := c.BindJSON(&job)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	job.Created = time.Now().Unix()

	if err := saveJob(job); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, job)
}

func saveJob(job producer.Job) error {
	b, err := json.Marshal(job)
	if err != nil {
		return err
	}

	configMap := kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092,localhost:9093,localhost:9094",
	}

	p, err := kafka.NewProducer(&configMap)
	if err != nil {
		return err
	}

	message := kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &producer.JobsTopic,
			Partition: kafka.PartitionAny,
		},
		Value: b,
	}

	if err := p.Produce(&message, nil); err != nil {
		return err
	}

	return nil
}
