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

func (api *Api) Execute() {
	apiGroup := api.app.Group("api")
	{
		newsGroup := apiGroup.Group("news")
		{
			newsGroup.GET("/:id", api.GetDetailNews)
			newsGroup.GET("/", api.GetNews)
			newsGroup.POST("/", api.SubmitNews)
			newsGroup.PUT("/:id", api.UpdateNews)
			newsGroup.DELETE("/:id", api.DeleteNews)
		}

		profileGroup := apiGroup.Group("profile")
		{
			profileGroup.GET("/:id", api.GetDetailProfile)
			profileGroup.GET("/", api.GetProfile)
			profileGroup.POST("/", api.SubmitProfile)
			profileGroup.PUT("/:id", api.UpdateProfile)
			profileGroup.DELETE("/:id", api.DeleteProfile)
		}
	}
}
