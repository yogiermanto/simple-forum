package memberships

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	membershipModel "simple-forum/internal/models/memberships"
	"simple-forum/internal/pkg/jwt"
	refresh_token "simple-forum/internal/pkg/refresh-token"
	"strconv"
	"time"
)

func (s *Service) Login(ctx context.Context, req membershipModel.LoginRequest) (membershipModel.LoginResponse, error) {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, "", 0)
	if err != nil {
		return membershipModel.LoginResponse{}, err
	}
	if user.ID == 0 {
		return membershipModel.LoginResponse{}, errors.New("email or password is not correct")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return membershipModel.LoginResponse{}, errors.New("email or password is not correct")
	}

	token, err := jwt.CreateToken(user.ID, user.Username, s.config.Service.SecretJWT)
	if err != nil {
		return membershipModel.LoginResponse{}, err
	}

	resp := membershipModel.LoginResponse{
		AccessToken: token,
	}

	now := time.Now()
	refreshTokenData, err := s.membershipRepo.GetRefreshToken(ctx, user.ID, now)
	if err != nil {
		return resp, err
	}

	//if already has refresh token will return login resp
	if refreshTokenData.ID != 0 {
		resp.RefreshToken = refreshTokenData.RefreshToken
		return resp, nil
	}

	userIDStr := strconv.FormatInt(user.ID, 10)
	refreshToken := refresh_token.GenerateRefreshToken()
	err = s.membershipRepo.InsertRefreshToken(ctx, membershipModel.RefreshTokenModel{
		UserID:       user.ID,
		RefreshToken: refreshToken,
		ExpiredAt:    time.Now().Add(10 * 24 * time.Hour),
		CreatedAt:    now,
		CreatedBy:    userIDStr,
		UpdatedAt:    now,
		UpdatedBy:    userIDStr,
	})
	if err != nil {
		return resp, err
	}

	resp.RefreshToken = refreshToken
	return resp, nil

}
