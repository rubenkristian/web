package api

import "github.com/gin-gonic/gin"

func (api *Api) GetNews(c *gin.Context) {
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

func (api *Api) GetDetailNews(c *gin.Context) {
	_, b := c.Params.Get("id")

	if !b {
		c.JSON(500, gin.H{
			"status": 500,
			"msg":    "id not found",
		})
		return
	}
}

func (api *Api) SubmitNews(c *gin.Context) {

}

func (api *Api) UpdateNews(c *gin.Context) {
	_, b := c.Params.Get("id")

	if !b {
		c.JSON(500, gin.H{
			"status": 500,
			"msg":    "id not found",
		})
		return
	}
}

// Delete Action
func (api *Api) DeleteNews(c *gin.Context) {
	_, b := c.Params.Get("id")

	if !b {
		c.JSON(500, gin.H{
			"status": 500,
			"msg":    "id not found",
		})
		return
	}
}
