package rabbitmq

import (
	"bytes"
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"sync"
)

type RabbitConsumer struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      *amqp.Queue
	Cache      *sync.Map
}

func NewConsumer() (*RabbitConsumer, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		"r6index", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return nil, err
	}

	return &RabbitConsumer{
		connection: conn,
		channel:    ch,
		queue:      &q,
		Cache:      &sync.Map{},
	}, nil
}

func (c *RabbitConsumer) Close() error {
	err := c.channel.Close()
	if err != nil {
		log.Println("error closing rabbit channel")
	}

	err = c.connection.Close()
	if err != nil {
		log.Println("error closing rabbit connection")
	}

	return nil
}

func (c *RabbitConsumer) write(b []byte) error {
	var session map[string]string

	err := json.NewDecoder(bytes.NewReader(b)).Decode(&session)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("writing to Sync Map", session)
	c.Cache.Store("session", session)
	return nil
}

func (c *RabbitConsumer) Consumer(ctx context.Context) {
	msgs, err := c.channel.Consume(
		c.queue.Name, // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)

	if err != nil {
		return
	}

	for {
		select {
		case <-ctx.Done():
			log.Println("Exiting Loop")
			return
		case msg := <-msgs:
			err := c.write(msg.Body)
			if err != nil {
				log.Println("could not write to session cache")
			}
		default:
			log.Println("no activity")
		}
	}
}
