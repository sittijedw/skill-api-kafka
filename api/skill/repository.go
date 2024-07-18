package skill

import "database/sql"

type SkillRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) SkillRepository {
	return SkillRepository{db: db}
}

func (repository *SkillRepository) findSkillByKey(key string) *sql.Row {
	sqlStatement := "SELECT key, name, description, logo, tags FROM skill where key=$1"
	return repository.db.QueryRow(sqlStatement, key)
}
