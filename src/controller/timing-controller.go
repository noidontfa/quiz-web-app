package controller

import (
	models "../model"
	"../service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TControl struct {
	TimingService service.TimingService
}

type TimingController interface {
	FindAllTimings(ctx *gin.Context)
	FindByIdTiming(ctx *gin.Context)
	UpdateTiming(ctx *gin.Context)
	DeleteTiming(ctx *gin.Context)
	SaveTiming(ctx *gin.Context)
}

func NewTimingController(serv service.TimingService) TimingController {
	return &TControl{TimingService:serv}
}


func (t *TControl) FindAllTimings(ctx *gin.Context) {
	timings, err := t.TimingService.FindAll()
	if err != nil {
		ctx.String(http.StatusInternalServerError,err.Error())
		return
	}
	ctx.JSON(http.StatusOK,timings)
}

func (t *TControl) FindByIdTiming(ctx *gin.Context) {
	id, err1 := strconv.ParseInt(ctx.Param("id"),0,0)
	if err1 != nil {
		ctx.String(http.StatusInternalServerError, err1.Error())
		return
	}
	timing, err := t.TimingService.FindById(uint(id))
	if err != nil {
		ctx.String(http.StatusInternalServerError,err.Error())
		return
	}
	ctx.JSON(http.StatusOK,timing)
}

func (t *TControl) UpdateTiming(ctx *gin.Context) {
	id, err1 := strconv.ParseInt(ctx.Param("id"),0,0)
	if err1 != nil {
		ctx.String(http.StatusInternalServerError, err1.Error())
		return
	}
	var timing models.Timing
	if err := ctx.ShouldBindJSON(&timing); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	timingResult, err := t.TimingService.Update(uint(id),&timing)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK,timingResult)
}

func (t *TControl) DeleteTiming(ctx *gin.Context) {
	id, err1 := strconv.ParseInt(ctx.Param("id"),0,0)
	if err1 != nil {
		ctx.String(http.StatusInternalServerError, err1.Error())
		return
	}
	err := t.TimingService.Delete(uint(id))
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.String(http.StatusOK, "Deleted")
}

func (t *TControl) SaveTiming(ctx *gin.Context) {
	var timing models.Timing
	if err := ctx.ShouldBindJSON(&timing); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	timingResult, err := t.TimingService.Save(&timing)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, timingResult)
}


