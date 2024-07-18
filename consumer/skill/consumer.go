package skill

import "github.com/IBM/sarama"

func ConsumeMessage(msg *sarama.ConsumerMessage, handler SkillHandler) {
	topic, key, value := msg.Topic, string(msg.Key), msg.Value

	if topic == "skill" {
		if key == "create" {
			handler.createHandler(value)
		}
	}
}
