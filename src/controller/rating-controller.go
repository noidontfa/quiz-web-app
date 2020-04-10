package controller

import (
	models "../model"
	"../service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RatingControl struct {
	RatingService service.RatingService
}

type RatingController interface {
	SaveRating(ctx *gin.Context)
}

func NewRatingController(serv service.RatingService) RatingController{
	return &RatingControl{RatingService:serv}
}

func (r *RatingControl) SaveRating(ctx *gin.Context) {
	var rating models.Rating
	if err := ctx.ShouldBindJSON(&rating); err != nil {
		ctx.String(http.StatusInternalServerError,err.Error())
		return
	}
	ratingResult, err := r.RatingService.Save(&rating)
	if err != nil {
		ctx.String(http.StatusInternalServerError,err.Error())
		return
	}
	ctx.JSON(http.StatusOK,ratingResult)
}


