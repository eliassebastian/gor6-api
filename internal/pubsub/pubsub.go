package pubsub

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

type Consumer struct {
	reader *kafka.Reader
}

func NewReader() *Consumer {
	return &Consumer{
		//TODO: secure connection
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:     []string{"localhost:9092"},
			Topic:       "ubisoft-topic",
			GroupID:     "ubisoft",
			MinBytes:    10e3, // 10KB
			MaxBytes:    10e6, // 10MB
			StartOffset: kafka.LastOffset,
		}),
	}
}

func (c *Consumer) Close() {
	err := c.reader.Close()
	if err != nil {
		log.Println(":::::WARNING: error closing kafka consumer reader")
	}
}

func (c *Consumer) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Println("Exiting Loop")
			return
		default:
			m, err := c.reader.FetchMessage(ctx)
			if err != nil {
				log.Println("Error Reading Kafka Message")
				continue
			}
			//TODO: send to ubisoft connection repo
			log.Println(m.Topic, m.Partition, m.Offset)
			//value := m.Value
			//log.Printf("message at topic/partition/offset %v / %v/ %v : %s\n", m.Topic, m.Partition, m.Offset, string(value))
		}
	}
}
