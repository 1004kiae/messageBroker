package messageBroker

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"os"
	"strconv"
	"time"
)

type Emqx struct {
	options *mqtt.ClientOptions

	qos      int
	retained bool
}

type EmqxClient struct {
	emqxClient mqtt.Client

	qos      int
	retained bool
}

func (e *Emqx) initialize(config *BrokerConfig) {
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

	e.options = opts
	e.qos = config.Qos
	e.retained = config.Retained
}

func (e *Emqx) createClient() (ClientInterface, error) {
	emqxClient := mqtt.NewClient(e.options)

	if token := emqxClient.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	client := new(EmqxClient)
	client.emqxClient = emqxClient
	client.qos = e.qos
	client.retained = e.retained

	return client, nil
}

func (ec *EmqxClient) Pub(topic string, message string) error {
	if token := ec.emqxClient.Publish(topic, byte(ec.qos), ec.retained, message); token.Error() != nil {
		return token.Error()
	}

	return nil
}

func (ec *EmqxClient) Sub(topic string, callbackFunc func([]byte)) error {
	if token := ec.emqxClient.Subscribe(topic, byte(ec.qos), func(client mqtt.Client, message mqtt.Message) {
		callbackFunc(message.Payload())
	}); token.Error() != nil {
		return token.Error()
	}

	return nil
}
