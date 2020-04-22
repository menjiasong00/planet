package srv

import (
	"rest/config"
	"rest/pb"
	"rest/service"
	"rest/util/rb"
)



var ConsumersMap = make(map[string][]interface{})


func makeConsumersMap() {

	switch *config.Server {
	case "bas":
		//死信消费者
		ConsumersMap["dlx.global.queue"] = []interface{}{"dlx.global.queue",&service.BasMqServer{},"DlxConsumer",&pb.DlxConsumerRequest{},rb.ReceiverConfig{1,1,true,false,false,true}}
		ConsumersMap["bas.msg.notify"] = []interface{}{"bas.msg.notify",&service.MsgServer{},"MsgDistributionCenter",&pb.MsgDistributionCenterRequest{},rb.ReceiverConfig{1,1,true,false,false,false}}

		break
	default:
		break
	}

}
