package main

import (
	_ "net/url"
)

func test(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{
		"message":"success",
	})
}



