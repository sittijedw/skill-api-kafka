package skill

import "database/sql"

type skillRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) skillRepository {
	return skillRepository{db: db}
}

func (repository *skillRepository) findSkillByKey(key string) *sql.Row {
	sqlStatement := "SELECT key, name, description, logo, tags FROM skill where key=$1"
	return repository.db.QueryRow(sqlStatement, key)
}

func (repository *skillRepository) findAll() (*sql.Rows, error) {
	sqlStatement := "SELECT key, name, description, logo, tags FROM skill"
	return repository.db.Query(sqlStatement)
}
