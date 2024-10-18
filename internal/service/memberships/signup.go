package memberships

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	membershipModel "simple-forum/internal/models/memberships"
	"time"
)

func (s *Service) SignUp(ctx context.Context, req membershipModel.SignUpRequest) error {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, req.Username, 0)
	if err != nil {
		return err
	}

	if user.ID != 0 {
		return errors.New("username or email already exists")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	now := time.Now()
	model := membershipModel.UserModel{
		Email:     req.Email,
		Password:  string(pass),
		Username:  req.Username,
		CreatedAt: now,
		CreatedBy: req.Email,
		UpdatedAt: now,
		UpdatedBy: req.Email,
	}

	if err = s.membershipRepo.CreateUser(ctx, model); err != nil {
		return err
	}

	return nil
}
