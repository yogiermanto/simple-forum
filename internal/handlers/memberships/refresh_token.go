package memberships

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-forum/internal/models/memberships"
)

func (h *Handler) RefreshToken(c *gin.Context) {
	ctx := c.Request.Context()

	var request memberships.RefreshTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetInt64("userID")
	resp, err := h.membershipSvc.ValidateRefreshToken(ctx, userID, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
