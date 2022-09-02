package routers

import (
  "github.com/gin-gonic/gin"
  "example.com/helloworld/controllers"
)

// func HelloWorld(c *gin.Context){
// 	name := c.Query("name")
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "HelloWorld "+ name,
// 	})
// }


func Routes() *gin.Engine{
  	router := gin.Default()
	router.GET("/HelloWorld", controllers.HelloWorld)
	router.GET("/ping", controllers.Pong)

	//   router.POST("/somePost", posting)
	//   router.PUT("/somePut", putting)
	//   router.DELETE("/someDelete", deleting)
	//   router.PATCH("/somePatch", patching)
	//   router.HEAD("/someHead", head)
	//   router.OPTIONS("/someOptions", options)

  return router
}