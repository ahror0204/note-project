package v1

import (
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/note_project/api/token"
	"github.com/note_project/config"
	"github.com/note_project/pkg/logger"
	"github.com/note_project/storage/repo"
)

type handlerV1 struct {
	log             logger.Logger
	cfg             config.Config
	postgresStorage repo.NoteRepositoryStorage
	redisStorage    repo.RedisRepositoryStorage
	jwtHandler      token.JWTHandler
}

type HandlerV1Config struct {
	Logger     logger.Logger
	Cfg        config.Config
	Postgres   repo.NoteRepositoryStorage
	Redis      repo.RedisRepositoryStorage
	jwtHandler token.JWTHandler
}

func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		log:             c.Logger,
		cfg:             c.Cfg,
		postgresStorage: c.Postgres,
		redisStorage:    c.Redis,
		jwtHandler:      c.jwtHandler,
	}
}

func CheckClaims(h *handlerV1, c *gin.Context) jwt.MapClaims {
	var (
		ErrUnauthorized = errors.New("unauthorized")
		authorization   JwtRequestModel
		claims          jwt.MapClaims
		err             error
	)

	authorization.Token = c.GetHeader("Authorization")
	if c.Request.Header.Get("Authorization") == "" {
		c.JSON(http.StatusUnauthorized, ErrUnauthorized)
		h.log.Error("Unauthorized request:", logger.Error(err))
	}

	h.jwtHandler.Token = authorization.Token
	claims, err = h.jwtHandler.ExtractClaims()
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrUnauthorized)
		h.log.Error("token is invalid:", logger.Error(err))
		return nil
	}
	return claims
}
