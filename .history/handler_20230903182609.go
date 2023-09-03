package main

import (

	_ "encoding/json"
	_ "errors"
	_ "fmt"
	"net/http"
	_ "net/url"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/julienschmidt/httprouter"
)

func test(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{
		"message":"success",
	})
}



