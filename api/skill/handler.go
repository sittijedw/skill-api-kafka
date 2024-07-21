package skill

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ResponseData interface {
	Skill | []Skill
}

type ResponseWithData[responseData Skill | []Skill] struct {
	Status string       `json:"status"`
	Data   responseData `json:"data"`
}

type SkillHandler struct {
	repository SkillRepository
}

func NewHandler(repository SkillRepository) SkillHandler {
	return SkillHandler{repository: repository}
}

func (handler *SkillHandler) GetByKeyHandler(ctx *gin.Context) {
	key := ctx.Param("key")

	row := handler.repository.findSkillByKey(key)

	skill, err := mapRowToSkill(row)

	if err != nil {
		responseError(ctx, "Skill not found", http.StatusNotFound)
		return
	}

	responseSuccessWithData(ctx, skill, http.StatusOK)
}

func (handler *SkillHandler) GetAllHandler(ctx *gin.Context) {
	rows, err := handler.repository.findAll()

	if err != nil {
		responseError(ctx, "Can't get all skills", http.StatusInternalServerError)
		return
	}

	var skills = make([]Skill, 0)
	for rows.Next() {
		var skill Skill

		err := rows.Scan(&skill.Key, &skill.Name, &skill.Description, &skill.Logo, pq.Array(&skill.Tags))

		if err != nil {
			responseError(ctx, "Can't scan row to skill struct", http.StatusInternalServerError)
			return
		}

		skills = append(skills, skill)
	}

	responseSuccessWithData(ctx, skills, http.StatusOK)
}

func createAndUpdateSkill(ctx *gin.Context, action string, respMessage string) {
	key := ctx.Param("key")

	var skill Skill

	if err := ctx.BindJSON(&skill); err != nil {
		responseError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	skillString, err := parseToString(skill)

	if err != nil {
		responseError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	if action == "create" {
		sendMessage(skillString, action)
	} else {
		sendMessage(skillString, action+"-"+key)
	}
	responseSuccess(ctx, respMessage, http.StatusOK)
}

func (handler *SkillHandler) CreateHandler(ctx *gin.Context) {
	createAndUpdateSkill(ctx, "create", "Creating Skill...")
}

func (handler *SkillHandler) UpdateByKeyHandler(ctx *gin.Context) {
	createAndUpdateSkill(ctx, "update", "Updating Skill...")
}

func (handler *SkillHandler) UpdateNameByKeyHandler(ctx *gin.Context) {
	createAndUpdateSkill(ctx, "update-name", "Updating Skill name...")
}

func (handler *SkillHandler) UpdateDescriptionByKeyHandler(ctx *gin.Context) {
	createAndUpdateSkill(ctx, "update-description", "Updating Skill description...")
}

func mapRowToSkill(row *sql.Row) (Skill, error) {
	var skill Skill
	err := row.Scan(&skill.Key, &skill.Name, &skill.Description, &skill.Logo, pq.Array(&skill.Tags))

	return skill, err
}

func responseError(ctx *gin.Context, message string, statusCode int) {
	response := Response{Status: "error", Message: message}
	ctx.JSON(statusCode, response)
}

func responseSuccess(ctx *gin.Context, message string, statusCode int) {
	response := Response{Status: "success", Message: message}
	ctx.JSON(statusCode, response)
}

func responseSuccessWithData[responseData ResponseData](ctx *gin.Context, data responseData, statusCode int) {
	response := ResponseWithData[responseData]{Status: "success", Data: data}
	ctx.JSON(statusCode, response)
}

func parseToString[responseData ResponseData](data responseData) (string, error) {
	productsJson, err := json.Marshal(data)

	if err != nil {
		return "", err
	}

	return string(productsJson), nil
}
