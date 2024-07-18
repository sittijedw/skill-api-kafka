package skill

import (
	"database/sql"
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

func mapRowToSkill(row *sql.Row) (Skill, error) {
	var skill Skill
	err := row.Scan(&skill.Key, &skill.Name, &skill.Description, &skill.Logo, pq.Array(&skill.Tags))

	return skill, err
}

func responseError(ctx *gin.Context, message string, statusCode int) {
	response := Response{Status: "error", Message: message}
	ctx.JSON(statusCode, response)
}

func responseSuccessWithData[responseData ResponseData](ctx *gin.Context, data responseData, statusCode int) {
	response := ResponseWithData[responseData]{Status: "success", Data: data}
	ctx.JSON(statusCode, response)
}
