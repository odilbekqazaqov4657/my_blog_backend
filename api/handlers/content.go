package handlers

import (
	"odilbekqazaqov4657/my_blog_backend/models"
	"odilbekqazaqov4657/my_blog_backend/pkg/helpers"

	"github.com/gin-gonic/gin"
	"github.com/saidamir98/udevs_pkg/logger"
)

func (h *handlers) CreateCategory(ctx *gin.Context) {

	var reqBody models.CreateCategoryReq

	err := ctx.BindJSON(&reqBody)
	if err != nil {
		h.log.Error("error on builting request body", logger.Error(err))
		return
	}

	category := &models.Category{}

	helpers.DataParser(reqBody, category)

	category, err = h.storage.GetContentRepo().CreateCategory(ctx, category)
	if err != nil {
		h.log.Error("error on creating new category", logger.Error(err))
		return
	}

	ctx.JSON(201, category)

}

func (h *handlers) GetCategoriesList(ctx *gin.Context) {

	page := ctx.Query("page")
	limit := ctx.Query("limit")

	categories, err := h.storage.GetContentRepo().GetCategories(ctx, helpers.GetPage(page), helpers.GetLimit(limit))

	if err != nil {
		h.log.Error("error on getting category list", logger.Error(err))
		return
	}
	ctx.JSON(200, categories)

}

func (h *handlers) GetCategory(ctx *gin.Context) {

	category := &models.Category{}

	id := ctx.Param("id")

	category, err := h.storage.GetContentRepo().GetCategory(ctx, id)
	if err != nil {
		h.log.Error("error on getting category", logger.Error(err))
		return
	}

	ctx.JSON(201, category)

}
