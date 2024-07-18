package skill

import (
	"encoding/json"
	"log"
)

type SkillHandler struct {
	repository SkillRepository
}

func NewHandler(repository SkillRepository) SkillHandler {
	return SkillHandler{repository: repository}
}

func (handler *SkillHandler) createHandler(data []byte) {
	skill, err := parseToSkill(data)

	if err != nil {
		log.Println("Error: Can't parse to skill struct")
		return
	}

	row := handler.repository.create(skill)

	err = row.Scan(&skill.Key)

	if err != nil {
		log.Println("Error: Skill already exists")
		return
	}
}

func parseToSkill(data []byte) (Skill, error) {
	var skill Skill
	err := json.Unmarshal([]byte(data), &skill)

	return skill, err
}
