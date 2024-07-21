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

func (repository *SkillRepository) updateByKey(skill Skill, skillKey string) *sql.Row {
	sqlStatement := "UPDATE skill SET name=$1, description=$2, logo=$3, tags=$4 WHERE key=$5 RETURNING key"
	return repository.db.QueryRow(sqlStatement, skill.Name, skill.Description, skill.Logo, pq.Array(skill.Tags), skillKey)
}

func (repository *SkillRepository) updateNameByKey(skill Skill, skillKey string) *sql.Row {
	sqlStatement := "UPDATE skill SET name=$1 WHERE key=$2 RETURNING key"
	return repository.db.QueryRow(sqlStatement, skill.Name, skillKey)
}
