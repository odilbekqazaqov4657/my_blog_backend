package api

import (
	"odilbekqazaqov4657/my_blog_backend/api/handlers"
	log "odilbekqazaqov4657/my_blog_backend/pkg/logger"
	"odilbekqazaqov4657/my_blog_backend/storage"
	"odilbekqazaqov4657/my_blog_backend/storage/redis"

	"github.com/gin-gonic/gin"
)

type Options struct {
	Storage storage.StorageI
	Log     log.Log
	Cache   redis.RedisRepoI
}

func Api(o Options) *gin.Engine {

	h := handlers.NewHandler(handlers.Handlers{Storage: o.Storage, Log: o.Log, Cache: o.Cache})

	engine := gin.Default()

	api := engine.Group("/api")

	api.GET("/ping", h.Ping)

	own := api.Group("/own") // faqat owner uchun ochiq endpointlar

	{
		//	own.POST("/log-in")
		//	own.POST("/log-out")
		own.POST("/category", h.CreateCategory)
		//	own.PUT("/category/:id")
		//	own.DELETE("/category/:id")
	}

	//	vw := api.Group("/vw") // faqat viewers uchun ochiq endpointlar
	//	{
	//		vw.POST("/log-out")
	//		vw.POST("/comment/:article_id")
	//	}
	//
	pb := api.Group("/pb") // hamma uchun ochiq endpointlar
	//	{
	pb.POST("/check_user", h.CheckUser)
	pb.POST("/check_otp", h.GetCategory)
	//		pb.POST("/sign_up")
	//		pb.POST("/log-in")
	pb.GET("/categories", h.GetCategoriesList)
	pb.GET("/categories/:id", h.GetCategory)
	//		pb.GET("/sub-categories")
	//		pb.GET("/sub-categories/:id")
	//		pb.GET("/articles")
	//		pb.GET("/articles/:id")
	//	}

	return engine

}
