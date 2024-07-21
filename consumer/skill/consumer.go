package skill

import (
	"strings"

	"github.com/IBM/sarama"
)

func ConsumeMessage(msg *sarama.ConsumerMessage, handler SkillHandler) {
	topic, key, value := msg.Topic, string(msg.Key), msg.Value

	splitKey := strings.Split(key, "-")

	var action string
	if len(splitKey) > 2 {
		action = splitKey[0] + "-" + splitKey[1]
	} else {
		action = splitKey[0]
	}

	skillKey := splitKey[len(splitKey)-1]

	if topic == "skill" {
		if action == "create" {
			handler.createHandler(value)
		} else if action == "update" {
			handler.updateByKeyHandler(value, skillKey)
		} else if action == "update-name" {
			handler.updateNameByKeyHandler(value, skillKey)
		} else if action == "update-description" {
			handler.updateDescriptionByKeyHandler(value, skillKey)
		} else if action == "update-logo" {
			handler.updateLogoByKeyHandler(value, skillKey)
		} else if action == "update-tags" {
			handler.updateTagsByKeyHandler(value, skillKey)
		}
	}
}
