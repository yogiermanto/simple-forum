package memberships

import (
	"context"
	"simple-forum/internal/configs"
	"simple-forum/internal/models/memberships"
	"time"
)

type membershipRepository interface {
	GetUser(ctx context.Context, email, username string, userID int64) (memberships.UserModel, error)
	CreateUser(ctx context.Context, model memberships.UserModel) error
	InsertRefreshToken(ctx context.Context, model memberships.RefreshTokenModel) error
	GetRefreshToken(ctx context.Context, userID int64, now time.Time) (memberships.RefreshTokenModel, error)
}

type Service struct {
	config         *configs.Config
	membershipRepo membershipRepository
}

func NewService(config *configs.Config, membershipRepo membershipRepository) *Service {
	return &Service{
		config:         config,
		membershipRepo: membershipRepo,
	}
}
