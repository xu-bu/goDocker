package main

import (

	_ "encoding/json"
	_ "errors"
	_ "fmt"
	"log"
	"math/big"
	"net/http"
	_ "net/url"
	"os"

	"goTest/interactTest/api"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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



