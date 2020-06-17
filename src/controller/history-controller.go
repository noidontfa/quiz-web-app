package controller

import (
	models "../model"
	"../service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type HistoryControl struct {
	HistoryService service.HistoryService
}

type HistoryController interface {
	SaveHistory(ctx *gin.Context)
	FindByIdHistory(ctx *gin.Context)
	FindByDateHistory(ctx *gin.Context)
}

func NewHistoryController(serv service.HistoryService) HistoryController {
	return &HistoryControl{HistoryService: serv}
}

func (h *HistoryControl) SaveHistory(ctx *gin.Context) {
	var history models.History
	if err := ctx.ShouldBindJSON(&history); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	historyResult, err := h.HistoryService.Save(&history)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, historyResult)
}

func (h *HistoryControl) FindByIdHistory(ctx *gin.Context) {
	quizId, _ := strconv.ParseInt(ctx.DefaultQuery("quizid", "0"), 0, 0)
	userId, _ := strconv.ParseInt(ctx.DefaultQuery("userid", "0"), 0, 0)
	if quizId > 0 && userId > 0 {
		ctx.String(http.StatusInternalServerError, "Id confict")
	} else if quizId == 0 && userId == 0 {
		ctx.String(http.StatusInternalServerError, "Not found")
	}
	if quizId > 0 {
		history, err := h.HistoryService.FindByQuizId(uint(quizId))
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, history)
		return
	}
	if userId > 0 {
		history, err := h.HistoryService.FindByUserId(uint(userId))
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, history)
		return
	}
}

func (h *HistoryControl) FindByDateHistory(ctx *gin.Context) {
	date := ctx.DefaultQuery("date", "")
	quizId, _ := strconv.ParseInt(ctx.DefaultQuery("quizid", "0"), 0, 0)
	if date == "" {
		ctx.String(http.StatusInternalServerError, "Date not founded")
		return
	}
	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Invalid Format")
		return
	}
	history, err := h.HistoryService.FindByDateId(date, uint(quizId))
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, history)
}
