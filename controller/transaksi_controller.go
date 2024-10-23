package controller

import (
	"net/http"

	"github.com/adieos/ets-pweb-be/dto"
	"github.com/adieos/ets-pweb-be/service"
	"github.com/adieos/ets-pweb-be/utils"
	"github.com/gin-gonic/gin"
)

type (
	TransaksiController interface {
		Create(ctx *gin.Context)
		GetAll(ctx *gin.Context)
		GetAllKomunal(ctx *gin.Context)
		GetById(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	transaksiController struct {
		transaksiService service.TransaksiService
	}
)

func NewTransaksiController(ts service.TransaksiService) TransaksiController {
	return &transaksiController{
		transaksiService: ts,
	}
}

func (c *transaksiController) Create(ctx *gin.Context) {
	user_id := ctx.MustGet("user_id").(string)
	var transaksi dto.CreateTransaksiRequest
	if err := ctx.ShouldBind(&transaksi); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.transaksiService.CreateTransaksi(user_id, transaksi)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_TRANSAKSI, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_TRANSAKSI, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *transaksiController) GetAll(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(string)
	role := ctx.MustGet("role").(string)
	result, err := c.transaksiService.GetAllTransaksi(userID, role)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_ALL_TRANSAKSI, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_ALL_TRANSAKSI, result)
	ctx.JSON(http.StatusOK, res)
}

// BELUM SELESAI !!!
func (c *transaksiController) GetAllKomunal(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(string)
	role := ctx.MustGet("role").(string)
	result, err := c.transaksiService.GetAllTransaksi(userID, role)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_ALL_TRANSAKSI, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_ALL_TRANSAKSI, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *transaksiController) GetById(ctx *gin.Context) {
	transaksiID := ctx.Param("id")
	result, err := c.transaksiService.GetDetailTransaksi(transaksiID)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_TRANSAKSI, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_TRANSAKSI, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *transaksiController) Update(ctx *gin.Context) {
	transaksiID := ctx.Param("id")
	var transaksi dto.UpdateTransaksiRequest
	if err := ctx.ShouldBind(&transaksi); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.transaksiService.UpdateTransaksi(transaksiID, transaksi)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_TRANSAKSI, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_TRANSAKSI, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *transaksiController) Delete(ctx *gin.Context) {
	transaksiID := ctx.Param("id")

	err := c.transaksiService.DeleteTransaksi(transaksiID)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_TRANSAKSI, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_TRANSAKSI, "transaksi deleted")
	ctx.JSON(http.StatusOK, res)
}
