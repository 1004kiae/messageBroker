package messageBroker

import (
	"fmt"
)

// noinspection ALL
const (
	BROKER_TYPE_EMQX = iota
	//BROKER_TYPE_REDIS
)

type BrokerConfig struct {
	BrokerType int

	Qos      int
	Retained bool

	Host string
	Port int

	UserName string
	Password string
}

type BrokerInterface interface {
	initialize(*BrokerConfig)
	createClient() (ClientInterface, error)
}

type ClientInterface interface {
	Pub(string, string) error
	Sub(string, func([]byte)) error
}

type MsgBroker struct {
	broker BrokerInterface
}

func (b *MsgBroker) Initialize(config *BrokerConfig) error {
	switch config.BrokerType {
	case BROKER_TYPE_EMQX:
		b.broker = new(Emqx)
		b.broker.initialize(config)
		break

	//case BROKER_TYPE_REDIS:
	//	b.broker = new(Redis)
	//	b.broker.initialize(config)
	//	break

	default:
		return fmt.Errorf("broker unknown type: %d", config.BrokerType)
	}

	return nil
}

func (b *MsgBroker) NewClient() (ClientInterface, error) {
	return b.broker.createClient()
}
