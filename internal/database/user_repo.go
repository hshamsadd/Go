package database

import (
	"Go/my-go-app/internal/models"
	"database/sql"
)

type PostgresUserRepository struct {
	DB *sql.DB
}

func (r *PostgresUserRepository) GetByID(id string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, email, name, created_at FROM users WHERE id = $1`
	
	err := r.DB.QueryRow(query, id).Scan(&user.ID, &user.Email, &user.Name, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}