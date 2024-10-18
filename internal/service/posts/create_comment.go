package posts

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	"simple-forum/internal/models/posts"
	"strconv"
	"time"
)

func (s *Service) CreateComment(ctx context.Context, postID, userID int64, request posts.CreateCommentRequest) error {
	post, err := s.postRepo.GetPostByID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get post by id from repository")
		return err
	}

	if post.ID == 0 {
		log.Error().Err(err).Msg("Post data is not found")
		return errors.New("post is not found")
	}

	now := time.Now()
	userIDStr := strconv.FormatInt(userID, 10)

	model := posts.CommentModel{
		PostID:         postID,
		UserID:         userID,
		CommentContent: request.CommentContent,
		CreatedAt:      now,
		CreatedBy:      userIDStr,
		UpdatedAt:      now,
		UpdatedBy:      userIDStr,
	}

	err = s.postRepo.CreateComment(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create comment from repository")
		return err
	}

	return nil
}
