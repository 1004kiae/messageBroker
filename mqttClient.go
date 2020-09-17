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
	broker broker.BrokerInterface
	client broker.ClientInterface
}

func (b *BrokerClient) Initialize(config *BrokerConfig) error {
	switch config.brokerType {
	case BROKER_TYPE_EMQX:
		b.broker = new(broker.Emqx)
		b.broker.SettingOptions(config)
		break
	default:
		return fmt.Errorf("broker unknown type: %d", config.brokerType)
	}

	return nil
}

func (b *BrokerClient) NewClient() (*broker.ClientInterface, error) {
	return b.broker.CreateClient()
}

func (b *BrokerClient) Pub(topic string, message string) error {
	return b.client.Pub(topic, message)
}

func (b *BrokerClient) Sub(topic string, callbackFunc func()) error {
	return b.client.Sub(topic, callbackFunc)
}
