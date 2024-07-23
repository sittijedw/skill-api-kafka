package skill

import (
	"database/sql"

	"github.com/lib/pq"
)

type skillRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) skillRepository {
	return skillRepository{db: db}
}

func (repository *skillRepository) create(skill skill) *sql.Row {
	sqlStatement := "INSERT INTO skill (key, name, description, logo, tags) VALUES ($1, $2, $3, $4, $5) RETURNING key"
	return repository.db.QueryRow(sqlStatement, skill.Key, skill.Name, skill.Description, skill.Logo, pq.Array(skill.Tags))
}

func (repository *skillRepository) updateByKey(skill skill, skillKey string) *sql.Row {
	sqlStatement := "UPDATE skill SET name=$1, description=$2, logo=$3, tags=$4 WHERE key=$5 RETURNING key"
	return repository.db.QueryRow(sqlStatement, skill.Name, skill.Description, skill.Logo, pq.Array(skill.Tags), skillKey)
}

func (repository *skillRepository) updateNameByKey(skill skill, skillKey string) *sql.Row {
	sqlStatement := "UPDATE skill SET name=$1 WHERE key=$2 RETURNING key"
	return repository.db.QueryRow(sqlStatement, skill.Name, skillKey)
}

func (repository *skillRepository) updateDescriptionByKey(skill skill, skillKey string) *sql.Row {
	sqlStatement := "UPDATE skill SET description=$1 WHERE key=$2 RETURNING key"
	return repository.db.QueryRow(sqlStatement, skill.Description, skillKey)
}

func (repository *skillRepository) updateLogoByKey(skill skill, skillKey string) *sql.Row {
	sqlStatement := "UPDATE skill SET logo=$1 WHERE key=$2 RETURNING key"
	return repository.db.QueryRow(sqlStatement, skill.Logo, skillKey)
}

func (repository *skillRepository) updateTagsByKey(skill skill, skillKey string) *sql.Row {
	sqlStatement := "UPDATE skill SET tags=$1 WHERE key=$2 RETURNING key"
	return repository.db.QueryRow(sqlStatement, pq.Array(skill.Tags), skillKey)
}

func (repository *skillRepository) deleteByKey(skillKey string) *sql.Row {
	sqlStatement := "DELETE FROM skill WHERE key=$1 RETURNING key"
	return repository.db.QueryRow(sqlStatement, skillKey)
}
