package posts

import (
	"context"
	"database/sql"
	"errors"
	"simple-forum/internal/models/posts"
)

func (r *Repository) GetUserActivity(ctx context.Context, model posts.UserActivityModel) (posts.UserActivityModel, error) {
	query := `SELECT id, post_id, user_id, is_liked, created_at, updated_at, created_by, updated_by FROM user_activities WHERE post_id = ? and user_id = ? LIMIT 1`

	var resp posts.UserActivityModel
	err := r.db.
		QueryRowContext(ctx, query, model.PostID, model.UserID).
		Scan(&resp.ID, &resp.PostID, &resp.UserID, &resp.IsLiked, &resp.CreatedAt, &resp.UpdatedAt, &resp.CreatedBy, &resp.UpdatedBy)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return resp, nil
		}
		return resp, err
	}

	return resp, nil
}

func (r *Repository) CreateUserActivity(ctx context.Context, model posts.UserActivityModel) error {
	query := `INSERT INTO user_activities (post_id, user_id, is_liked, created_at, updated_at, created_by, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, model.PostID, model.UserID, model.IsLiked, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateUserActivity(ctx context.Context, model posts.UserActivityModel) error {
	query := `UPDATE user_activities SET post_id = ?, user_id = ?, is_liked = ?, created_at = ?, updated_at = ?, created_by = ?, updated_by = ? where post_id = ? and user_id = ?`

	_, err := r.db.ExecContext(ctx, query, model.PostID, model.UserID, model.IsLiked, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy, model.PostID, model.UserID)
	if err != nil {
		return err
	}

	return nil
}
