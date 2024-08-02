package handlers

import (
	log "odilbekqazaqov4657/my_blog_backend/pkg/logger"
	"odilbekqazaqov4657/my_blog_backend/storage"
	"odilbekqazaqov4657/my_blog_backend/storage/redis"

	"github.com/gin-gonic/gin"
)

type handlers struct {
	storage storage.StorageI
	log     log.Log
	cache   redis.RedisRepoI
}

type Handlers struct {
	Storage storage.StorageI
	Log     log.Log
	Cache   redis.RedisRepoI
}

func NewHandler(h Handlers) handlers {
	return handlers{h.Storage, h.Log, h.Cache}
}

func (h *handlers) Ping(ctx *gin.Context) {
	ctx.JSON(200, map[string]string{"message": "pong"})
}
