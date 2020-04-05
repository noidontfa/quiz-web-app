package controller

import (
	models "../model"
	"../service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


type LangControl struct {
	LanguageService service.LanguageService
}

type LanguageController interface {
	FindAllLanguages(ctx *gin.Context)
	FindByIdLanguage(ctx *gin.Context)
	UpdateLanguage(ctx *gin.Context)
	DeleteLanguage(ctx *gin.Context)
	SaveLanguage(ctx *gin.Context)
}

func NewLanguageController(sevc service.LanguageService) LanguageController{
	return &LangControl{LanguageService:sevc}
}

func (l *LangControl) FindAllLanguages(ctx *gin.Context) {
	languages, err := l.LanguageService.FindAll()
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK,languages)
}

func (l *LangControl) FindByIdLanguage(ctx *gin.Context) {
	id, err1 := strconv.ParseInt(ctx.Param("id"),0,0)
	if err1 != nil {
		ctx.String(http.StatusInternalServerError, err1.Error())
		return
	}
	language, err := l.LanguageService.FindById(uint(id))
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK,language)
}

func (l *LangControl) UpdateLanguage(ctx *gin.Context) {
	id, err1 := strconv.ParseInt(ctx.Param("id"),0,0)
	if err1 != nil {
		ctx.String(http.StatusInternalServerError, err1.Error())
		return
	}
	var language models.Language
	if err := ctx.ShouldBindJSON(&language); err != nil {
		ctx.String(http.StatusNoContent,err.Error())
		return
	}

	languageResult, err := l.LanguageService.Update(uint(id),&language)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, languageResult)
}

func (l *LangControl) DeleteLanguage(ctx *gin.Context) {
	id, err1 := strconv.ParseInt(ctx.Param("id"),0,0)
	if err1 != nil {
		ctx.String(http.StatusInternalServerError, err1.Error())
		return
	}
	err := l.LanguageService.Delete(uint(id))
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.String(http.StatusOK, "Deleted")
}

func (l *LangControl) SaveLanguage(ctx *gin.Context) {
	var language models.Language
	if err := ctx.ShouldBindJSON(&language); err != nil {
		ctx.String(http.StatusInternalServerError,err.Error())
		return
	}
	languageResult, err := l.LanguageService.Save(&language)
	if err != nil {
		ctx.String(http.StatusInternalServerError,err.Error())
		return
	}
	ctx.JSON(http.StatusOK,languageResult)
}

