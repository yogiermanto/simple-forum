package memberships

import (
	"context"
	"errors"
	"simple-forum/internal/models/memberships"
	"simple-forum/internal/pkg/jwt"
	"time"
)

func (s *Service) ValidateRefreshToken(ctx context.Context, userID int64, request memberships.RefreshTokenRequest) (resp memberships.RefreshTokenResponse, err error) {
	refreshTokenData, err := s.membershipRepo.GetRefreshToken(ctx, userID, time.Now())
	if err != nil {
		return resp, err
	}
	if refreshTokenData.ID == 0 {
		return resp, errors.New("refresh token has expired")
	}
	if refreshTokenData.RefreshToken != request.RefreshToken {
		return resp, errors.New("invalid refresh token")
	}

	user, err := s.membershipRepo.GetUser(ctx, "", "", userID)
	if err != nil {
		return resp, err
	}
	if user.ID == 0 {
		return resp, errors.New("email or password is not correct")
	}

	token, err := jwt.CreateToken(user.ID, user.Username, s.config.Service.SecretJWT)
	if err != nil {
		return resp, err
	}

	resp.AccessToken = token
	return resp, err
}
