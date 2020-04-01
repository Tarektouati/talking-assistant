# Talking Assistant 🤖

A Worker that consumes a given queue and makes your home assistant say something interesting.

## Getting started

#### Docker :

```bash
docker run -e BROKER_CONNECTION_STRING="XXX" -e BROKER_QUEUE="XXX" -e BROKER="XXX" -e ASSISTANT_HOST="XXX" -e ASSISTANT_LANG="XXX" -e ASSISTANT_ACCENT="XXX" tarektouati/talking-assistant
```

#### Go :

```bash
 go build -o talking-assistant cmd/talking-assistant/main.go
 ./talking-assistant
```

## Configuration

### Supported brokers

- RabbitMQ
- MQTT (TODO)

### Setup your environment variables

All environment variables are **required** and there's no default configuration

| Name                     | Description                                       |
| ------------------------ | ------------------------------------------------- |
| BROKER_CONNECTION_STRING | Broker connectionstring                           |
| BROKER_QUEUE             | Queue name                                        |
| BROKER                   | Broker type (check the list of supported brokers) |
| ASSISTANT_HOST           | Hostname of your vocal assistant                  |
| ASSISTANT_LANG           | Language for your vocal assistant                 |
| ASSISTANT_ACCENT         | Accent for your vocal assistant                   |
