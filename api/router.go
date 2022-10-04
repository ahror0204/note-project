package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	v1 "github.com/note_project/api/handlers/v1"
	"github.com/note_project/config"
	"github.com/note_project/pkg/logger"
	repo "github.com/note_project/storage/repo"

	_ "github.com/note_project/api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Option struct {
	DB        *sqlx.DB
	Conf      config.Config
	Logger    logger.Logger
	Users     repo.UserRepositoryStorage
	Postgres  repo.NoteRepositoryStorage
	RedisRepo repo.RedisRepositoryStorage
}

func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Cfg:         option.Conf,
		Logger:      option.Logger,
		UserStorage: option.Users,
		Postgres:    option.Postgres,
		Redis:       option.RedisRepo,
	})

	api := router.Group("/v1")

	api.POST("/notes/", handlerV1.SetNoteWithTTL)
	api.POST("/createnote/", handlerV1.CreateNote)
	api.PUT("/updatenote/", handlerV1.UpdateNote)
	api.DELETE("/deletenote/:id", handlerV1.DeleteNote)

	api.POST("/create_user/", handlerV1.CreateUser)
	api.POST("/update_user/", handlerV1.UpdateUser)
	api.DELETE("/delete_user/:id", handlerV1.DeleteUser)
	api.POST("/register/", handlerV1.RegisterUser)
	api.POST("/users/verify_user/", handlerV1.VerifyUser)

	url := ginSwagger.URL("swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
