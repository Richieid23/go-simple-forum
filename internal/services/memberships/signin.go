package memberships

import (
	"context"
	"errors"
	"github.com/Richieid23/simple-forum/internal/models/memberships"
	"github.com/Richieid23/simple-forum/pkg/jwt"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) SignIn(ctx context.Context, req memberships.SignInRequest) (string, error) {
	user, err := s.membershipRepos.GetUser(ctx, "", req.Email)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user")
		return "", err
	}

	if user == nil {
		return "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	token, err := jwt.CreateToken(user.ID, user.Username, s.config.Jwt.SecretKey)
	if err != nil {
		return "", err
	}

	return token, nil

}
