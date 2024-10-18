package posts

import (
	"context"
	"github.com/gin-gonic/gin"
	"simple-forum/internal/middleware"
	"simple-forum/internal/models/posts"
)

type postService interface {
	CreatePost(ctx context.Context, userID int64, req posts.CreatePostRequest) error
	CreateComment(ctx context.Context, postID, userID int64, request posts.CreateCommentRequest) error
	UpsertUserActivity(ctx context.Context, postID, userID int64, request posts.UserActivityRequest) error
	GetALlPost(ctx context.Context, pageSize, pageIndex int) (posts.GetAllPostResponse, error)
}

type Handler struct {
	*gin.Engine
	postSvc postService
}

func NewHandler(api *gin.Engine, postSvc postService) *Handler {
	return &Handler{
		api,
		postSvc,
	}
}

func (h *Handler) RegisterRoutes() {
	r := h.Group("/posts")
	r.Use(middleware.AuthMiddleware())
	r.GET("", h.GetAllPost)
	r.POST("", h.CreatePost)

	r.POST("/:id/comments", h.CreateComment)
	r.PUT("/:id/user-activity", h.UpsertUserActivity)
}
