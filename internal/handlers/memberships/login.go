package memberships

import (
	"github.com/gin-gonic/gin"
	"net/http"
	membershipModel "simple-forum/internal/models/memberships"
)

func (h *Handler) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var request membershipModel.LoginRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.membershipSvc.Login(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
