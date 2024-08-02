package redis

import (
	"context"
	log "odilbekqazaqov4657/my_blog_backend/pkg/logger"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/saidamir98/udevs_pkg/logger"
)

type RedisRepoI interface {
	Exists(ctx context.Context, key string) (bool, error)
	Set(ctx context.Context, key, value string, exp int) error
	Get(ctx context.Context, key string) (any, error)
	GetDelete(ctx context.Context, key string) (any, error)
	Delete(ctx context.Context, key string) (any, error)
}

type redisRepo struct {
	cli redis.Client
	log log.Log
}

func NewRedisRepo(cli redis.Client, log log.Log) RedisRepoI {
	return &redisRepo{
		cli: cli,
		log: log,
	}
}

// Exists
func (r *redisRepo) Exists(ctx context.Context, key string) (bool, error) {

	defer r.cli.Close()

	isExists, err := r.cli.Do(ctx, "EXISTS", key).Result()

	if err != nil {
		r.log.Error("error on checking exists", logger.Error(err))
		return false, err
	}

	exists, _ := isExists.(int)

	return exists == 1, nil

}

// Set
func (r *redisRepo) Set(ctx context.Context, key, value string, exp int) error {

	_, err := r.cli.SetEX(ctx, key, value, time.Second*time.Duration(exp)).Result()

	if err != nil {
		r.log.Error("error on setting to cache", logger.Error(err))
		return err
	}

	return nil
}

// Get
func (r *redisRepo) Get(ctx context.Context, key string) (any, error) {
	return nil, nil
}

// Get delete
func (r *redisRepo) GetDelete(ctx context.Context, key string) (any, error) {
	return nil, nil
}

// Delete
func (r *redisRepo) Delete(ctx context.Context, key string) (any, error) {
	return nil, nil
}
