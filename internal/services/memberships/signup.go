package memberships

import (
	"context"
	"errors"
	"github.com/Richieid23/simple-forum/internal/models/memberships"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (s *service) SignUp(ctx context.Context, req memberships.SignUpRequest) error {
	user, err := s.membershipRepos.GetUser(ctx, req.Username, req.Email)
	if err != nil {
		return err
	}

	if user != nil {
		return errors.New("username or email already exists")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	now := time.Now()
	data := memberships.UserModel{
		Username:  req.Username,
		Email:     req.Email,
		Password:  string(pass),
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: req.Username,
		UpdatedBy: req.Username,
	}

	return s.membershipRepos.CreateUser(ctx, data)
}
