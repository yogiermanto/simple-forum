package memberships

import (
	"context"
	"database/sql"
	"errors"
	"simple-forum/internal/models/memberships"
	"time"
)

func (r *Repository) InsertRefreshToken(ctx context.Context, model memberships.RefreshTokenModel) error {
	query := `INSERT INTO refresh_tokens (user_id, refresh_token, expired_at, created_at, updated_at, created_by, updated_by)
			 VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, model.UserID, model.RefreshToken, model.ExpiredAt, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetRefreshToken(ctx context.Context, userID int64, now time.Time) (memberships.RefreshTokenModel, error) {
	query := `SELECT id, user_id, refresh_token, expired_at, created_at, updated_at, created_by, updated_by 
			FROM refresh_tokens 
			WHERE user_id = ? and expired_at >= ?
			order by created_at DESC
			LIMIT 1`

	var resp memberships.RefreshTokenModel
	err := r.db.
		QueryRowContext(ctx, query, userID, now).
		Scan(&resp.ID, &resp.UserID, &resp.RefreshToken, &resp.ExpiredAt, &resp.CreatedAt, &resp.UpdatedAt, &resp.CreatedBy, &resp.UpdatedBy)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return resp, nil
		}
		return resp, err
	}

	return resp, nil
}
