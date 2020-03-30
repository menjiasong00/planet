package topic

type Topic interface {
	GetQueueName() string
	GetRoutingKeys() []string
	Execute(routingKey string, data interface{}) error
}