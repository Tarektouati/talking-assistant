package amqp

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
	"github.com/tarektouati/talking-assistant/pkg/broker"
)

// Client for amqp broker
type Client struct {
	ConnectionString string
	QueueName        string
}

//getEnvWithError func returns error if the env is not found
func getEnvWithError(env string) (string, error) {
	envValue, found := os.LookupEnv(env)
	if !found {
		return "", fmt.Errorf("%s env not found", env)
	}
	return envValue, nil
}

//NewClient creates a new Client for amqp broker
func NewClient() (broker.Broker, error) {
	connectionString, err := getEnvWithError("BROKER_CONNECTION_STRING")
	if err != nil {
		return nil, err
	}
	queue, err := getEnvWithError("BROKER_QUEUE")
	if err != nil {
		return nil, err
	}

	client := &Client{
		ConnectionString: connectionString,
		QueueName:        queue,
	}
	return client, nil
}

func createConnection(connectionSting string) (*amqp.Connection, error) {
	return amqp.Dial(connectionSting)
}

func createQueue(conn *amqp.Connection, queueName string) (*amqp.Channel, amqp.Queue, error) {
	channel, err := conn.Channel()
	if err != nil {
		return nil, amqp.Queue{}, err
	}

	queue, err := channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return nil, amqp.Queue{}, err
	}

	err = channel.Qos(1, 0, false)
	if err != nil {
		return nil, amqp.Queue{}, err
	}
	return channel, queue, nil
}

//Consume consumes broker's topic and executes the callback on recived message
func (c *Client) Consume(onMessage func(message string)) error {
	conn, err := createConnection(c.ConnectionString)
	if err != nil {
		return err
	}
	defer conn.Close()

	amqpChannel, queue, err := createQueue(conn, c.QueueName)
	if err != nil {
		return err
	}
	defer amqpChannel.Close()

	messageChannel, err := amqpChannel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	stopChan := make(chan bool)
	go func() {
		log.Printf("Consumer ready")
		for d := range messageChannel {
			log.Printf("Received a message: %s", string(d.Body))
			onMessage(string(d.Body))
			if err := d.Ack(false); err != nil {
				log.Printf("Error acknowledging message : %s", err)
			} else {
				log.Printf("Acknowledged message")
			}
		}
	}()
	<-stopChan
	return nil
}
