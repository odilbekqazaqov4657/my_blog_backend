package db

import (
	"context"
	"odilbekqazaqov4657/my_blog_backend/config"
	log "odilbekqazaqov4657/my_blog_backend/pkg/logger"
	"strconv"

	"github.com/saidamir98/udevs_pkg/logger"

	"github.com/go-redis/redis/v8"
)

func RedisAddr(host string, port int) string {
	return host + ":" + strconv.Itoa(port)
}

func ConnRedis(log log.Log, ctx context.Context, cfg config.RedisConfig) (*redis.Client, error) {

	redisCli := redis.NewClient(&redis.Options{

		Addr: RedisAddr(cfg.Host, cfg.Port),
		DB:   cfg.DbIndex,
	})

	_, err := redisCli.Ping(ctx).Result()

	if err != nil {
		log.Error("error on connecting to redis ", logger.Error(err))
		return nil, err
	}

	return redisCli, nil

}
