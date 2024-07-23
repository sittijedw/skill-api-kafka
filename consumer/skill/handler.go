package skill

import (
	"database/sql"
	"encoding/json"
	"log"
)

type skillHandler struct {
	repository skillRepository
}

func NewHandler(repository skillRepository) skillHandler {
	return skillHandler{repository: repository}
}

func (handler *skillHandler) createAndUpdateSkill(data []byte, action string, skillKey string, errMessage string) {
	skill, err := parseToSkill(data)

	if err != nil {
		log.Println("Error: Can't parse to skill struct :", err)
		return
	}

	var row *sql.Row
	if skillKey == "" && action == "create" {
		row = handler.repository.create(skill)
	} else if skillKey != "" {
		if action == "update" {
			row = handler.repository.updateByKey(skill, skillKey)
		} else if action == "update-name" {
			row = handler.repository.updateNameByKey(skill, skillKey)
		} else if action == "update-description" {
			row = handler.repository.updateDescriptionByKey(skill, skillKey)
		} else if action == "update-logo" {
			row = handler.repository.updateLogoByKey(skill, skillKey)
		} else if action == "update-tags" {
			row = handler.repository.updateTagsByKey(skill, skillKey)
		}
	}

	err = row.Scan(&skill.Key)

	if err != nil {
		log.Println(errMessage, ":", err)
		return
	}
}

func (handler *skillHandler) deleteSkill(skillKey string, errMessage string) {
	var skill skill
	row := handler.repository.deleteByKey(skillKey)
	err := row.Scan(&skill.Key)

	if err != nil {
		log.Println(errMessage, ":", err)
		return
	}
}

func (handler *skillHandler) createHandler(data []byte) {
	handler.createAndUpdateSkill(data, "create", "", "Error: Skill already exists")
}

func (handler *skillHandler) updateByKeyHandler(data []byte, skillKey string) {
	handler.createAndUpdateSkill(data, "update", skillKey, "Error: Unable to update skill")
}

func (handler *skillHandler) updateNameByKeyHandler(data []byte, skillKey string) {
	handler.createAndUpdateSkill(data, "update-name", skillKey, "Error: Unable to update skill name")
}

func (handler *skillHandler) updateDescriptionByKeyHandler(data []byte, skillKey string) {
	handler.createAndUpdateSkill(data, "update-description", skillKey, "Error: Unable to update skill description")
}

func (handler *skillHandler) updateLogoByKeyHandler(data []byte, skillKey string) {
	handler.createAndUpdateSkill(data, "update-logo", skillKey, "Error: Unable to update skill logo")
}

func (handler *skillHandler) updateTagsByKeyHandler(data []byte, skillKey string) {
	handler.createAndUpdateSkill(data, "update-tags", skillKey, "Error: Unable to update skill tags")
}

func (handler *skillHandler) deleteByKeyHandler(skillKey string) {
	handler.deleteSkill(skillKey, "Error: Unable to delete skill")
}

func parseToSkill(data []byte) (skill, error) {
	var skill skill
	err := json.Unmarshal([]byte(data), &skill)

	return skill, err
}
