package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context){
	name := c.Query("name")
	var message string = " this is your first api in Golang"
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello "+ name + message,
	})
}

func Pong(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
