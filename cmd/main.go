package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"simple-forum/internal/configs"
	"simple-forum/internal/handlers/memberships"
	"simple-forum/internal/handlers/posts"
	"simple-forum/internal/pkg/internalsql"
	membershipsRepository "simple-forum/internal/repository/memberships"
	postsRepository "simple-forum/internal/repository/posts"
	membershipsService "simple-forum/internal/service/memberships"
	postsService "simple-forum/internal/service/posts"
)

func main() {
	r := gin.Default()

	err := configs.Init()
	if err != nil {
		panic("failed to initialize configs: " + err.Error())
	}
	conf := configs.Get()

	db, err := internalsql.Connect(conf.Database.DataSourceName)
	if err != nil {
		panic("failed to initialize database connection: " + err.Error())
	}

	log.Error().Err(errors.New("test config")).Msg("test config dsn" + conf.Database.DataSourceName)
	log.Error().Err(errors.New("test config")).Msg("test config secret" + conf.Service.SecretJWT)

	membershipsRepo := membershipsRepository.NewRepository(db)
	postsRepo := postsRepository.NewRepository(db)

	membershipsSvc := membershipsService.NewService(conf, membershipsRepo)
	postsSvc := postsService.NewService(conf, postsRepo)

	membershipsHandler := memberships.NewHandler(r, membershipsSvc)
	membershipsHandler.RegisterRoute()

	postsHandler := posts.NewHandler(r, postsSvc)
	postsHandler.RegisterRoutes()

	if err := r.Run(":8080"); err != nil {
		panic("failed to run router: " + err.Error())
	}
}
