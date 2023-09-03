package main

import (
	"github.com/gin-gonic/gin"
	""
)

func test(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{
		"message":"success",
	})
}



