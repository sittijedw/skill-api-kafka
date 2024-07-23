package skill

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type responseData interface {
	skill | []skill
}

type responseWithData[respData skill | []skill] struct {
	Status string   `json:"status"`
	Data   respData `json:"data"`
}

type skillHandler struct {
	repository skillRepository
}

func NewHandler(repository skillRepository) skillHandler {
	return skillHandler{repository: repository}
}

func (handler *skillHandler) GetByKeyHandler(ctx *gin.Context) {
	key := ctx.Param("key")

	row := handler.repository.findSkillByKey(key)

	skill, err := mapRowToSkill(row)

	if err != nil {
		responseError(ctx, "Skill not found", http.StatusNotFound)
		return
	}

	responseSuccessWithData(ctx, skill, http.StatusOK)
}

func (handler *skillHandler) GetAllHandler(ctx *gin.Context) {
	rows, err := handler.repository.findAll()

	if err != nil {
		responseError(ctx, "Can't get all skills", http.StatusInternalServerError)
		return
	}

	var skills = make([]skill, 0)
	for rows.Next() {
		var skill skill

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

	var skill skill

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

func deleteSkill(ctx *gin.Context, action string, respMessage string) {
	key := ctx.Param("key")

	sendMessage("", action+"-"+key)
	responseSuccess(ctx, respMessage, http.StatusOK)
}

func (handler *skillHandler) CreateHandler(ctx *gin.Context) {
	createAndUpdateSkill(ctx, "create", "Creating Skill...")
}

func (handler *skillHandler) UpdateByKeyHandler(ctx *gin.Context) {
	createAndUpdateSkill(ctx, "update", "Updating Skill...")
}

func (handler *skillHandler) UpdateNameByKeyHandler(ctx *gin.Context) {
	createAndUpdateSkill(ctx, "update-name", "Updating Skill name...")
}

func (handler *skillHandler) UpdateDescriptionByKeyHandler(ctx *gin.Context) {
	createAndUpdateSkill(ctx, "update-description", "Updating Skill description...")
}

func (handler *skillHandler) UpdateLogoByKeyHandler(ctx *gin.Context) {
	createAndUpdateSkill(ctx, "update-logo", "Updating Skill logo...")
}

func (handler *skillHandler) UpdateTagsByKeyHandler(ctx *gin.Context) {
	createAndUpdateSkill(ctx, "update-tags", "Updating Skill tags...")
}

func (handler *skillHandler) DeleteByKeyHandler(ctx *gin.Context) {
	deleteSkill(ctx, "delete", "Deleting Skill...")
}

func mapRowToSkill(row *sql.Row) (skill, error) {
	var skill skill
	err := row.Scan(&skill.Key, &skill.Name, &skill.Description, &skill.Logo, pq.Array(&skill.Tags))

	return skill, err
}

func responseError(ctx *gin.Context, message string, statusCode int) {
	response := response{Status: "error", Message: message}
	ctx.JSON(statusCode, response)
}

func responseSuccess(ctx *gin.Context, message string, statusCode int) {
	response := response{Status: "success", Message: message}
	ctx.JSON(statusCode, response)
}

func responseSuccessWithData[respData responseData](ctx *gin.Context, data respData, statusCode int) {
	response := responseWithData[respData]{Status: "success", Data: data}
	ctx.JSON(statusCode, response)
}

func parseToString[respData responseData](data respData) (string, error) {
	productsJson, err := json.Marshal(data)

	if err != nil {
		return "", err
	}

	return string(productsJson), nil
}
