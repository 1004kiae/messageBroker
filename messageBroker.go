package messageBroker

import "fmt"

// noinspection ALL
const (
	BROKER_TYPE_EMQX = iota
	//BROKER_TYPE_REDIS
)

type BrokerInterface interface {
	initialize(*BrokerConfig)
	createClient() (ClientInterface, error)
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

type MsgBroker struct {
	broker BrokerInterface
}

func (b *MsgBroker) Initialize(config *BrokerConfig) error {
	switch config.brokerType {
	case BROKER_TYPE_EMQX:
		b.broker = new(Emqx)
		b.broker.initialize(config)
		break

	//case BROKER_TYPE_REDIS:
	//	b.broker = new(Redis)
	//	b.broker.initialize(config)
	//	break

	default:
		return fmt.Errorf("broker unknown type: %d", config.brokerType)
	}

	return nil
}

func (b *MsgBroker) NewClient() (ClientInterface, error) {
	return b.broker.createClient()
}
