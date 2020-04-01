package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/evalphobia/google-home-client-go/googlehome"
	"github.com/tarektouati/talking-assistant/pkg/broker"
	"github.com/tarektouati/talking-assistant/pkg/broker/amqp"
)

//getEnvWithError func returns error if the env is not found
func getEnvWithError(env string) (string, error) {
	envValue, found := os.LookupEnv(env)
	if !found {
		return "", fmt.Errorf("%s env not found", env)
	}
	return envValue, nil
}

func getBrokerInstance(name string) (broker.Broker, error) {
	switch name {
	case "amqp":
		return amqp.NewClient()
	default:
		return nil, errors.New("Unknown broken")
	}
}

func createAssitantClient() (*googlehome.Client, error) {
	host, err := getEnvWithError("ASSISTANT_HOST")
	if err != nil {
		return nil, err
	}
	lang, err := getEnvWithError("ASSISTANT_LANG")
	if err != nil {
		return nil, err
	}
	accent, err := getEnvWithError("ASSISTANT_ACCENT")
	if err != nil {
		return nil, err
	}
	return googlehome.NewClientWithConfig(googlehome.Config{
		Hostname: host,
		Lang:     lang,
		Accent:   accent,
	})
}

func startApp() error {
	home, err := createAssitantClient()
	if err != nil {
		return err
	}
	log.Printf("Connected to Assistant")

	givenBroker, err := getEnvWithError("BROKER")
	if err != nil {
		return err
	}
	brokerinstance, err := getBrokerInstance(givenBroker)
	if err != nil {
		return err
	}
	if err = brokerinstance.Consume(func(message string) { home.Notify(message) }); err != nil {
		return err
	}
	return home.QuitApp()
}

func main() {
	if err := startApp(); err != nil {
		log.Fatal(err)
	}
}
