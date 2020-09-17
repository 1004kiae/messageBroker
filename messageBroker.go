package messageBroker

import (
	"fmt"
	"test/messageBroker/broker"
)

// noinspection ALL
const (
	BROKER_TYPE_EMQX = iota
	//BROKER_TYPE_REDIS
)

type BrokerInterface interface {
	Initialize(*BrokerConfig)
	CreateClient() (*ClientInterface, error)
}

type ClientInterface interface {
	Pub(string, string) error
	Sub(string, func()) error
}

type BrokerConfig struct {
	brokerType int

	Qos      int
	Retained bool

	Host string
	Port int

	UserName string
	Password string
}

type MessageBroker struct {
	broker *BrokerInterface
}

func (b *MessageBroker) Initialize(config *BrokerConfig) error {
	switch config.brokerType {
	case BROKER_TYPE_EMQX:
		br := new(broker.Emqx)
		br.Initialize(config)

		b.broker = new(broker.Emqx)
		break

	//case BROKER_TYPE_REDIS:
	//	b.broker = new(broker.Redis)
	//	b.broker.initialize(config)
	//	break

	default:
		return fmt.Errorf("broker unknown type: %d", config.brokerType)
	}

	return nil
}

func (b *MessageBroker) NewClient() (*ClientInterface, error) {
	//return b.broker.createClient()
}
