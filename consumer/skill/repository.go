package skill

import (
	"database/sql"

	"github.com/lib/pq"
)

type SkillRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) SkillRepository {
	return SkillRepository{db: db}
}

func (repository *SkillRepository) create(skill Skill) *sql.Row {
	sqlStatement := "INSERT INTO skill (key, name, description, logo, tags) VALUES ($1, $2, $3, $4, $5) RETURNING key"
	return repository.db.QueryRow(sqlStatement, skill.Key, skill.Name, skill.Description, skill.Logo, pq.Array(skill.Tags))
}
