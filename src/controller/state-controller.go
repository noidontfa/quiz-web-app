package controller

import (
	"../service"
	"github.com/gin-gonic/gin"
	"net/http"
)
type StateControl struct {
	StateService service.StateService
}


type StateController interface {
	FindAllState(ctx *gin.Context)
}

func NewStateController(serv service.StateService) StateController {
	return &StateControl{
		StateService:serv,
	}
}

func (s *StateControl) FindAllState(ctx *gin.Context) {
	states, err := s.StateService.FindAll()
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK,states)
}
