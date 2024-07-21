package skill

import (
	"database/sql"
	"encoding/json"
	"log"
)

type SkillHandler struct {
	repository SkillRepository
}

func NewHandler(repository SkillRepository) SkillHandler {
	return SkillHandler{repository: repository}
}

func (handler *SkillHandler) createAndUpdateSkill(data []byte, action string, skillKey string, errMessage string) {
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

func (handler *SkillHandler) deleteSkill(skillKey string, errMessage string) {
	var skill Skill
	row := handler.repository.deleteByKey(skillKey)
	err := row.Scan(&skill.Key)

	if err != nil {
		log.Println(errMessage, ":", err)
		return
	}
}

func (handler *SkillHandler) createHandler(data []byte) {
	handler.createAndUpdateSkill(data, "create", "", "Error: Skill already exists")
}

func (handler *SkillHandler) updateByKeyHandler(data []byte, skillKey string) {
	handler.createAndUpdateSkill(data, "update", skillKey, "Error: Unable to update skill")
}

func (handler *SkillHandler) updateNameByKeyHandler(data []byte, skillKey string) {
	handler.createAndUpdateSkill(data, "update-name", skillKey, "Error: Unable to update skill name")
}

func (handler *SkillHandler) updateDescriptionByKeyHandler(data []byte, skillKey string) {
	handler.createAndUpdateSkill(data, "update-description", skillKey, "Error: Unable to update skill description")
}

func (handler *SkillHandler) updateLogoByKeyHandler(data []byte, skillKey string) {
	handler.createAndUpdateSkill(data, "update-logo", skillKey, "Error: Unable to update skill logo")
}

func (handler *SkillHandler) updateTagsByKeyHandler(data []byte, skillKey string) {
	handler.createAndUpdateSkill(data, "update-tags", skillKey, "Error: Unable to update skill tags")
}

func (handler *SkillHandler) deleteByKeyHandler(skillKey string) {
	handler.deleteSkill(skillKey, "Error: Unable to delete skill")
}

func parseToSkill(data []byte) (Skill, error) {
	var skill Skill
	err := json.Unmarshal([]byte(data), &skill)

	return skill, err
}
