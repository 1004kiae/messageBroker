package broker

type Redis struct {
	//options *redis.Option
}

//
//type RedisClient struct {
//	redisClient // 레디스객체
//}
//
//func (r *Redis) Initialize(config MessageBrokerConfig) {
//	// r.options 에 redis 설정 저장
//}
//
//func (r *Redis) CreateClient() (*MessageBrokerClientInterface, error) {
//	redisClient := redis에서 connection 한 객체
//
//	client := new(RedisClient)
//	client.redisClient = redisClient
//
//	return client, nil
//}
//
//func (rc *RedisClient) Pub(topic string, message string) error {
//
//}
//
//func (rc *RedisClient) Sub(topic string, callbackFunc func()) error {
//
//}
