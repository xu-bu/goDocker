package main

import (

	_ "encoding/json"
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



