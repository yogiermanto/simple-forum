package posts

import (
	"context"
	"database/sql"
	"errors"
	"golang.org/x/sync/errgroup"
	"simple-forum/internal/models/posts"
)

func (r *Repository) CreatePost(ctx context.Context, model posts.PostModel) error {
	query := `INSERT INTO posts(user_id, post_title, post_content, post_hashtags, created_at, updated_at, created_by, updated_by)
				VALUES (?, ?, ?, ?, ?, ?, ?, ?)
			`

	_, err := r.db.ExecContext(ctx, query, model.UserID, model.PostTitle, model.PostContent, model.PostHashtags, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetPostByID(ctx context.Context, postID int64) (posts.PostModel, error) {
	query := `SELECT id, user_id, post_title, post_content, post_hashtags, created_at, updated_at, created_by, updated_by from posts where id = ? limit 1`

	var resp posts.PostModel
	err := r.db.
		QueryRowContext(ctx, query, postID).
		Scan(&resp.ID, &resp.UserID, &resp.PostTitle, &resp.PostContent, &resp.PostHashtags, &resp.CreatedAt, &resp.UpdatedAt, &resp.CreatedBy, &resp.UpdatedBy)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return posts.PostModel{}, nil
		}
		return posts.PostModel{}, err
	}

	return resp, nil
}

func (r *Repository) GetAllPost(ctx context.Context, limit, offset int) (posts.GetAllPostResponse, error) {
	errg, _ := errgroup.WithContext(ctx)

	var resp posts.GetAllPostResponse
	resp.Data = make([]posts.Post, 0)
	errg.Go(func() error {
		query := `SELECT p.id, p.user_id, u.username, p.post_title, p.post_content, p.post_hashtags 
				FROM posts p 
				JOIN users u ON p.user_id = u.id
				ORDER BY p.created_at DESC
				LIMIT ? OFFSET ?`

		rows, err := r.db.QueryContext(ctx, query, limit, offset)
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var (
				model    posts.PostModel
				username string
			)

			err := rows.Scan(&model.ID, &model.UserID, &username, &model.PostTitle, &model.PostContent, &model.PostHashtags)
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					return nil
				}
				return err
			}

			resp.Data = append(resp.Data, posts.Post{
				ID:           model.ID,
				UserID:       model.UserID,
				Username:     username,
				PostTitle:    model.PostTitle,
				PostContent:  model.PostContent,
				PostHashtags: model.PostHashtags,
			})
		}
		return nil
	})

	errg.Go(func() error {
		query := `SELECT count(p.id)
				FROM posts p 
				JOIN users u ON p.user_id = u.id`

		err := r.db.QueryRowContext(ctx, query).Scan(&resp.Total)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil
			}
			return err
		}
		return nil
	})

	err := errg.Wait()
	if err != nil {
		return resp, err
	}

	return resp, nil
}
