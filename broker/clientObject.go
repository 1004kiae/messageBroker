package broker

import "test/messageBroker"

// noinspection ALL
type BrokerInterface interface {
	SettingOptions(config *messageBroker.BrokerConfig)
	CreateClient() (*ClientInterface, error)
}

// noinspection ALL
type ClientInterface interface {
	Pub(string, string) error
	Sub(string, func()) error
}
