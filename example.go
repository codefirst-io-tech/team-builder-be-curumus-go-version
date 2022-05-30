package main

import "github.com/gin-gonic/gin"

func main() {
	baseApi := "/api/name"
	var names []string
	r := gin.Default()

	r.GET(baseApi, func(context *gin.Context) {
		context.JSON(200, names)
	})

	r.POST(baseApi+"/:name", func(context *gin.Context) {
		name := context.Params.ByName("name")
		names = append(names, name)
		context.JSON(200, names)
	})

	r.DELETE(baseApi+"/all", func(context *gin.Context) {
		names = []string{}
	})

	r.Run()
}
