package main

import (
	"context"
	"fmt"
	"odilbekqazaqov4657/my_blog_backend/api"
	"odilbekqazaqov4657/my_blog_backend/config"
	"odilbekqazaqov4657/my_blog_backend/pkg/db"
	log "odilbekqazaqov4657/my_blog_backend/pkg/logger"
	"odilbekqazaqov4657/my_blog_backend/storage"
	"odilbekqazaqov4657/my_blog_backend/storage/redis"

	"github.com/saidamir98/udevs_pkg/logger"
)

var ctx = context.Background()

func main() {

	cfg := config.Load() // configuraysiya yasab beradi

	log := log.NewLogger(cfg.GeneralConfig)

	pgxPool, err := db.ConnDb(cfg.PgConfig)

	if err != nil {
		log.Error("error on connecting to postgres !", logger.Error(err))
		return
	}

	log.Debug("successfully connected to postgres")

	fmt.Println(pgxPool)

	redisCli, err := db.ConnRedis(log, ctx, cfg.RedisConfig)

	if err != nil {
		log.Error("error on connecting to redis !", logger.Error(err))
		return
	}

	log.Debug("successfully connected to redis")

	cache := redis.NewRedisRepo(*redisCli, log)

	storage := storage.NewStorage(pgxPool, log)

	engine := api.Api(api.Options{
		Storage: storage,
		Log:     log,
		Cache:   cache,
	})

	log.Debug("server is running on ", logger.String("port", cfg.GeneralConfig.HTTPPort))

	engine.Run(cfg.GeneralConfig.HTTPPort)
}
