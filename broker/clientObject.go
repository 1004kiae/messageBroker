package broker

import "test/messageBroker"

// noinspection ALL
type BrokerClientInterface interface {
	SettingOptions(config *messageBroker.BrokerConfig)
	CreateClient() (*BrokerClientInterface, error)
	Pub(string, string) error
	Sub(string, func()) error
}
