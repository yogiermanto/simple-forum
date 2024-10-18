package posts

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	"simple-forum/internal/models/posts"
	"strconv"
	"time"
)

func (s *Service) UpsertUserActivity(ctx context.Context, postID, userID int64, request posts.UserActivityRequest) error {
	now := time.Now()
	userIDStr := strconv.FormatInt(userID, 10)
	model := posts.UserActivityModel{
		UserID:    userID,
		PostID:    postID,
		IsLiked:   request.IsLiked,
		CreatedAt: now,
		CreatedBy: userIDStr,
		UpdatedAt: now,
		UpdatedBy: userIDStr,
	}

	userActivity, err := s.postRepo.GetUserActivity(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user activity from repository")
		return err
	}

	if userActivity.ID == 0 {
		//create user activity

		//error if first time unlike but user_activities data is not found
		if !request.IsLiked {
			return errors.New("can't unlike post for the first time")
		}
		err = s.postRepo.CreateUserActivity(ctx, model)
		if err != nil {
			log.Error().Err(err).Msg("Failed to create user activity from repository")
			return err
		}
	} else {
		//update user activity
		err = s.postRepo.UpdateUserActivity(ctx, model)
		if err != nil {
			log.Error().Err(err).Msg("Failed to update user activity from repository")
			return err
		}
	}

	return nil
}
