package memberships

import (
	"context"
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"simple-forum/internal/models/memberships"
)

func (r *Repository) GetUser(ctx context.Context, email, username string, userID int64) (memberships.UserModel, error) {
	query := `SELECT id, email, password, username, created_at, created_by, updated_at, updated_by FROM users WHERE email = ? OR username = ? OR id = ?`

	var resp memberships.UserModel
	err := r.db.
		QueryRowContext(ctx, query, email, username, userID).
		Scan(&resp.ID, &resp.Email, &resp.Password, &resp.Username, &resp.CreatedAt, &resp.CreatedBy, &resp.UpdatedAt, &resp.UpdatedBy)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return memberships.UserModel{}, nil
		}
		return memberships.UserModel{}, err
	}

	return resp, nil
}

func (r *Repository) CreateUser(ctx context.Context, model memberships.UserModel) error {
	query := `INSERT INTO users (email, password, username, created_at, updated_at, created_by, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, model.Email, model.Password, model.Username, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)
	if err != nil {
		return err
	}

	return nil
}
