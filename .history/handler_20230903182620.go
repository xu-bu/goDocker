package main

import (
	_ "errors"
	_ "fmt"
	"net/http"
	_ "net/url"
)

func test(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{
		"message":"success",
	})
}



