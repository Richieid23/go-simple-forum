package memberships

import (
	"context"
	"github.com/Richieid23/simple-forum/internal/configs"
	"github.com/Richieid23/simple-forum/internal/models/memberships"
)

type membershipRepository interface {
	GetUser(ctx context.Context, username, email string) (*memberships.UserModel, error)
	CreateUser(ctx context.Context, user memberships.UserModel) error
}

type service struct {
	membershipRepos membershipRepository
	config          *configs.Config
}

func NewService(membershipRepos membershipRepository, configs *configs.Config) *service {
	return &service{membershipRepos: membershipRepos, config: configs}
}
