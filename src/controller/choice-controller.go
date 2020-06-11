package controller

import (
	"../service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


type ChoiceControl struct {
	ChoiceService service.ChoiceService
}

type ChoiceController interface {
	DeleteChoices(ctx *gin.Context)
}

func NewChoiceController(serv service.ChoiceService) ChoiceController {
	return &ChoiceControl{ChoiceService:serv}
}

func (c *ChoiceControl) DeleteChoices(ctx *gin.Context) {
	choiceId, err1 := strconv.ParseInt(ctx.Param("id"),0,0)
	if err1 != nil {
		ctx.String(http.StatusInternalServerError, err1.Error())
		return
	}
	err := c.ChoiceService.Delete(uint(choiceId))
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.String(http.StatusOK,"Deleted")
}

