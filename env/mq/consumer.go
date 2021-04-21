package server

import (
	"planet/pb"
	"planet/pkg/grbmq"
	"planet/service"
)

func ConfigConsumer(servFlag string)  []grbmq.ConsumerSetting {
	var ConsumerSettings []grbmq.ConsumerSetting
	switch servFlag {
	case "Bas":
		//基础
		ConsumerSettings = []grbmq.ConsumerSetting{
			//死信消费者
			{
				QueueName:"dlx.global.queue",
				RoutingKey:"dlx.global.queue",
				Service:&service.TestServer{},
				Controller:"DlxConsumer",
				Request:&pb.DlxConsumerRequest{},
				Workers:2,
				Config : grbmq.ReceiverConfig{1, 1, true, false, false, true},
			},
		}
		break
		case "Pro":
		//商品

		break
		case "Usr":
		//用户

		break
		default:

		break
	}

	return ConsumerSettings

}
