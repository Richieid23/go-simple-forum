package memberships

import (
	"context"
	"github.com/Richieid23/simple-forum/internal/models/memberships"
	"github.com/gin-gonic/gin"
)

type membershipService interface {
	SignUp(ctx context.Context, req memberships.SignUpRequest) error
	SignIn(ctx context.Context, req memberships.SignInRequest) (string, error)
}

type Handler struct {
	membershipService membershipService
	*gin.Engine
}

func NewHandler(api *gin.Engine, membershipService membershipService) *Handler {
	return &Handler{
		Engine:            api,
		membershipService: membershipService,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("/memberships")
	route.GET("/ping", h.Ping)
	route.POST("/sign-up", h.SignUp)
	route.POST("/sign-in", h.SignIn)
}
