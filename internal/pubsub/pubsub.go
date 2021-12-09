package pubsub

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
	"sync"
)

type Consumer struct {
	reader *kafka.Reader
	Cache  *sync.Map
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
		Cache: &sync.Map{},
	}
}

func (c *Consumer) Close() {
	err := c.reader.Close()
	if err != nil {
		log.Println(":::::WARNING: error closing kafka consumer reader")
	}
}

func (c *Consumer) Run(ctx context.Context) {
	log.Println("Starting Kafka Consumer")
	for {
		select {
		case <-ctx.Done():
			log.Println("Exiting Loop")
			return
		default:
			m, err := c.reader.FetchMessage(ctx)
			log.Println("Receiving Kafka Message")
			if err != nil {
				log.Println("Error Reading Kafka Message")
				continue
			}
			//TODO: send to ubisoft connection repo
			log.Println(m.Topic, m.Partition, m.Offset)
			we := c.write(m.Value)
			if we != nil {
				log.Println("Error Writing to Sync Map")
				continue
			}
			//value := m.Value
			//log.Printf("message at topic/partition/offset %v / %v/ %v : %s\n", m.Topic, m.Partition, m.Offset, string(value))
		}
	}
}

func (c *Consumer) write(b []byte) error {
	var session map[string]string
	log.Println(b)
	//var session models.Session
	err := json.NewDecoder(bytes.NewReader(b)).Decode(&session)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("SUCCESS writing to Sync Map", session)
	c.Cache.Store("session", session)
	return nil
}
