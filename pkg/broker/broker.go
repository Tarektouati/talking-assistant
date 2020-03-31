package broker

type Broker interface {
	Consume(func(message string)) error
}
