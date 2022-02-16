package api

import "github.com/gin-gonic/gin"

type Api struct {
	app *gin.Engine
}

func New(_app *gin.Engine) *Api {
	return &Api{
		app: _app,
	}
}
