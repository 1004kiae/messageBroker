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

type BrokerConfig struct {
	brokerType int

	Qos      int
	Retained bool

	Host string
	Port int

	UserName string
	Password string
}

type BrokerClient struct {
	client broker.BrokerClientInterface
}

func (b *BrokerClient) Initialize(config *BrokerConfig) error {
	switch config.brokerType {
	case BROKER_TYPE_EMQX:
		b.client = new(broker.Emqx)
		b.client.SettingOptions(config)
		break
	default:
		return fmt.Errorf("broker unknown type: %d", config.brokerType)
	}

	return nil
}

func (b *BrokerClient) NewClient() (*broker.BrokerClientInterface, error) {
	return b.client.CreateClient()
}
