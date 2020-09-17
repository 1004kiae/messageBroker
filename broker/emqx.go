package broker

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"os"
	"strconv"
	"test/messageBroker"
	"time"
)

// noinspection ALL
type Emqx struct {
	options *mqtt.ClientOptions

	qos      int
	retained bool
	client   mqtt.Client
}

func (b *Emqx) SettingOptions(config *messageBroker.BrokerConfig) {
	opts := mqtt.NewClientOptions()

	// The full URL of the MQTT server to connect to
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", config.Host, config.Port))

	// A username&password to authenticate to the MQTT server
	opts.SetUsername(config.UserName)
	opts.SetPassword(config.Password)

	// A clientID for the connection
	hostname, _ := os.Hostname()
	opts.SetClientID(fmt.Sprintf("%s_%s", hostname, strconv.Itoa(time.Now().Second())))

	//opts.SetCleanSession(true)

	b.options = opts
	b.qos = config.Qos
	b.retained = config.Retained
}

func (b *Emqx) CreateClient() (*BrokerClientInterface, error) {
	//opts := mqtt.NewClientOptions()
	//
	//// The full URL of the MQTT server to connect to
	//opts.AddBroker(fmt.Sprintf("tcp://%s:%d", mqttServer.Host, mqttServer.Port))
	//
	//// A username&password to authenticate to the MQTT server
	//opts.SetUsername(mqttServer.UserName)
	//opts.SetPassword(mqttServer.Password)
	//
	//// A clientID for the connection
	//hostname, _ := os.Hostname()
	//opts.SetClientID(fmt.Sprintf("%s_%s", hostname, strconv.Itoa(time.Now().Second())))
	//
	//client := mqtt.NewClient(opts)
	//
	//if token := client.Connect(); token.Wait() && token.Error() != nil {
	//	return token.Error()
	//}

	return nil, nil
}

func (b *Emqx) Pub(topic string, message string) error {
	if token := b.client.Publish(topic, byte(b.qos), b.retained, message); token.Error() != nil {
		return token.Error()
	}

	return nil
}

func (b *Emqx) Sub(topic string, callbackFunc func()) error {
	//	if token := b.client.Subscribe(topic, byte(b.qos), callback); token.Wait() && token.Error() != nil {
	//		return token.Error()
	//	}

	return nil
}

//func (b *BrokerClient) Sub(topic string, callback mqtt.MessageHandler) error {
//	if token := b.client.Subscribe(topic, byte(b.qos), callback); token.Wait() && token.Error() != nil {
//		return token.Error()
//	}
//
//	return nil
//}
