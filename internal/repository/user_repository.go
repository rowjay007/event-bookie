package repository

import (
    "context"
    "database/sql"
    "time"

    "github.com/rowjay007/event-bookie/internal/models"
)

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
    query := `
        INSERT INTO users (username, email, password, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5)
    `
    _, err := r.db.ExecContext(ctx, query, user.Username, user.Email, user.Password, time.Now(), time.Now())
    if err != nil {
        return err
    }
    return nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int) (*models.User, error) {
    query := `
        SELECT id, username, email, password, created_at, updated_at
        FROM users
        WHERE id = $1
    `
    row := r.db.QueryRowContext(ctx, query, id)
    user := &models.User{}
    err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
    if err != nil {
        return nil, err
    }
    return user, nil
}

// Implement other methods like UpdateUser, DeleteUser, GetUsers, etc.
