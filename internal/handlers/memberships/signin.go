package memberships

import (
	"github.com/Richieid23/simple-forum/internal/models/memberships"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) SignIn(c *gin.Context) {
	ctx := c.Request.Context()

	var request memberships.SignInRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := h.membershipService.SignIn(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := memberships.SignInResponse{accessToken}

	c.JSON(http.StatusOK, response)
}
