package memberships

import (
	"context"
	"github.com/gin-gonic/gin"
	"simple-forum/internal/middleware"
	membershipModel "simple-forum/internal/models/memberships"
)

type membershipService interface {
	SignUp(ctx context.Context, req membershipModel.SignUpRequest) error
	Login(ctx context.Context, req membershipModel.LoginRequest) (membershipModel.LoginResponse, error)
	ValidateRefreshToken(ctx context.Context, userID int64, request membershipModel.RefreshTokenRequest) (resp membershipModel.RefreshTokenResponse, err error)
}

type Handler struct {
	*gin.Engine
	membershipSvc membershipService
}

func NewHandler(api *gin.Engine, membershipSvc membershipService) *Handler {
	return &Handler{
		api,
		membershipSvc,
	}
}

func (h *Handler) RegisterRoute() {
	r := h.Group("memberships")
	r.GET("/ping", h.Ping)
	r.POST("/signup", h.SignUp)
	r.POST("/login", h.Login)

	r2 := h.Group("memberships")
	r2.Use(middleware.AuthRefreshTokenMiddleware())
	r2.POST("/refresh", h.RefreshToken)
}
