package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rubenkristian/myweb/api"
)

func main() {
	r := gin.Default()

	apiHandler := api.New(r)
	apiHandler.Execute()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
