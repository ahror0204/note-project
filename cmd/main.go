package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/note_project/api"
	"github.com/note_project/config"
	db "github.com/note_project/pkg/db"
	"github.com/note_project/pkg/logger"
	"github.com/note_project/storage/postgres"
	rds "github.com/note_project/storage/redis"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "test_task")
	defer logger.Cleanup(log)

	connDB, err := db.ConnectDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}
	defer connDB.Close()
	pool := redis.Pool{

		MaxIdle: 80,

		MaxActive: 12000,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort))
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}

	redisRepo := rds.NewRedisRepo(&pool)

	server := api.New(api.Option{

		Conf:      cfg,
		Logger:    log,
		Users: postgres.NewUserRepo(connDB),
		Postgres:  postgres.NewnoteRepo(connDB),
		RedisRepo: redisRepo,
	})

	if err := server.Run(cfg.HTTPport); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
	}
}
