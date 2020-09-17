package messageBroker

import (
	"fmt"
	"github.com/go-redis/redis"
)

type Redis struct {
	options *redis.Options
}

type RedisClient struct {
	redisClient *redis.Client
}

func (e *Redis) initialize(config *BrokerConfig) {
	redisAddr := fmt.Sprintf("%s:%d", config.Host, config.Port)

	opts := redis.Options{
		Addr:     redisAddr,
		Password: config.Password,
		//DialTimeout:  time.Second * time.Duration(config.Timeout),
		//ReadTimeout:  time.Second * time.Duration(config.Timeout),
		//WriteTimeout: time.Second * time.Duration(config.Timeout),
		//PoolSize:     config.PoolSize,
	}
	e.options = &opts
}

func (e *Redis) createClient() (ClientInterface, error) {
	redisClient := redis.NewClient(e.options)

	if err := redisClient.Ping().Err(); err != nil {
		return nil, err
	}

	client := new(RedisClient)
	client.redisClient = redisClient

	return client, nil
}

func (ec *RedisClient) Pub(topic string, message string) error {
	//if token := ec.redisClient.Publish(topic, byte(ec.qos), ec.retained, message); token.Error() != nil {
	//	return token.Error()
	//}

	return nil
}

func (ec *RedisClient) Sub(topic string, callbackFunc func([]byte)) error {
	//if token := ec.redisClient.Subscribe(topic, byte(ec.qos), func(client mqtt.Client, message mqtt.Message) {
	//	callbackFunc(message.Payload())
	//}); token.Error() != nil {
	//	return token.Error()
	//}

	return nil
}
