package posts

import (
	"context"
	"github.com/rs/zerolog/log"
	"simple-forum/internal/models/posts"
	"strconv"
	"strings"
	"time"
)

func (s *Service) CreatePost(ctx context.Context, userID int64, req posts.CreatePostRequest) error {
	postHashtags := strings.Join(req.PostHashtags, ",")
	userIDStr := strconv.FormatInt(userID, 10)

	now := time.Now()
	model := posts.PostModel{
		UserID:       userID,
		PostTitle:    req.PostTitle,
		PostContent:  req.PostContent,
		PostHashtags: postHashtags,
		CreatedAt:    now,
		CreatedBy:    userIDStr,
		UpdatedAt:    now,
		UpdatedBy:    userIDStr,
	}

	err := s.postRepo.CreatePost(ctx, model)
	if err != nil {
		log.Err(err).Msg("error create post to repository")
		return err
	}

	return nil
}
