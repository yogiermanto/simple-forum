package posts

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-forum/internal/models/posts"
	"strconv"
)

func (h *Handler) UpsertUserActivity(c *gin.Context) {
	ctx := c.Request.Context()

	var request posts.UserActivityRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	postIDStr := c.Param("id")
	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetInt64("userID")
	err = h.postSvc.UpsertUserActivity(ctx, postID, userID, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
