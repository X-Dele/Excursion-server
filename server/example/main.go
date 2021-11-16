package main

import "github.com/gin-gonic/gin"

//func main() {
//	r := gin.Default()
//	r.GET("/ping", func(c *gin.Context) {
//		c.JSON(200, gin.H{
//			"message": "pong",
//		})
//	})
//	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
//}

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()
	posting := func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "somePost",
		})
	}
	router.GET("/someGet", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "someGet",
		})
	})
	router.POST("/somePost", posting)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
	// router.Run(":3000") for a hard coded port
}
