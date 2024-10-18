package memberships

import (
	"github.com/gin-gonic/gin"
	"net/http"
	membershipModel "simple-forum/internal/models/memberships"
)

func (h *Handler) SignUp(c *gin.Context) {
	ctx := c.Request.Context()

	var request membershipModel.SignUpRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.membershipSvc.SignUp(ctx, request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "membership created"})
	return
}
