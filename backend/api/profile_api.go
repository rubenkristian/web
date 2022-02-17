package api

import "github.com/gin-gonic/gin"

func (web *Api) GetProfile(c *gin.Context) {
	var (
		pagination *PaginationQuery
	)

	err := c.ShouldBindQuery(&pagination)

	if err != nil {
		c.JSON(500, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}
}

func (web *Api) GetDetailProfile(c *gin.Context) {
	_, b := c.Params.Get("id")

	if !b {
		c.JSON(500, gin.H{
			"status": 500,
			"msg":    "id not found",
		})
		return
	}
}

func (web *Api) SubmitProfile(c *gin.Context) {

}

func (web *Api) UpdateProfile(c *gin.Context) {
	_, b := c.Params.Get("id")

	if !b {
		c.JSON(500, gin.H{
			"status": 500,
			"msg":    "id not found",
		})
		return
	}
}

func (web *Api) DeleteProfile(c *gin.Context) {
	_, b := c.Params.Get("id")

	if !b {
		c.JSON(500, gin.H{
			"status": 500,
			"msg":    "id not found",
		})
		return
	}
}
