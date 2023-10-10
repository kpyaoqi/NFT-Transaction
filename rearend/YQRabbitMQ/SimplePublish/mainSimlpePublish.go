package SendMQMessage

import (
	"yqnft/YQRabbitMQ/RabbitMQ"
)

var rabbitmq = RabbitMQ.NewRabbitMQSimple("KP")

func ChangeNftWorkByIdMessage(worksId string) {
	rabbitmq.PublishSimple("ChangeNftWorkByIdMessage_" + worksId)
}

func ChangeNftCollectionByIdMessage(collectionId string) {
	rabbitmq.PublishSimple("ChangeNftCollectionByIdMessage_" + collectionId)
}
