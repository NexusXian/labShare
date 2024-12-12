package main

import (
	"Test/myfunc"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	//制定静态文件路径
	r.Static("/s", "static")
	r.GET("/1", myfunc.Test1)
	r.GET("/2", myfunc.Test2)
	r.GET("/3", myfunc.Test3)
	r.GET("/4", myfunc.Test4)
	r.GET("/5", myfunc.Test5)
	r.GET("/6", myfunc.Test6)
	r.GET("/7", myfunc.Test7)

	r.Run(":8080")

}
