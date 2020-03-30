package topic

type TodoTopic struct {}

func (c TodoTopic) GetQueueName() string {
	return "todo_topic"
}
func (c TodoTopic) GetRoutingKeys() []string{
	return []string{"xxx","xxx2"}
}

func (c TodoTopic) Execute(routingKey string, data interface{}) error {
	
	return nil
}