package handlers

import (
	"encoding/json"
	"odilbekqazaqov4657/my_blog_backend/models"
	"odilbekqazaqov4657/my_blog_backend/pkg/mail"

	"github.com/gin-gonic/gin"
	"github.com/saidamir98/udevs_pkg/logger"
)

func (h *handlers) CheckUser(ctx *gin.Context) {

	reqBody := models.CheckViewer

	gin.Bind(reqBody)

	isExists, err := h.storage.GetCommonRepo().CheckIsExists(ctx, &models.Common{
		TableName:  "viewers",
		ColumnName: "gmail",
		ExpValue:   reqBody.Gmail,
	})

	if err != nil {
		h.log.Error("error on checking viewer", logger.Error(err))
		return
	}

	if isExists {

		ctx.JSON(201, models.CheckExistsResp{
			IsExists: isExists,
			Status:   "log-in",
		})
		return
	}

	otp := models.OtpData{
		Gmail: reqBody.Gmail,
		Otp:   mail.GenerateOtp(6),
	}

	otpDataB, err := json.Marshal(otp)
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	err = h.cache.Set(ctx, reqBody.Gmail, string(otpDataB), 60)
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	err = mail.SendMail([]string{reqBody.Gmail}, otp.Otp)
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, models.CheckExistsResp{
		IsExists: isExists,
		Status:   "register",
	})

}

func (h *handlers) CheckOTP(ctx *gin.Context) {

	var reqBody models.OtpData

	err := ctx.Bind(&reqBody)
	if err != nil {
		return
	}

	gmail := reqBody.Gmail

}
