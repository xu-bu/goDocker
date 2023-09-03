package main

import (
	"context"
	"crypto/ecdsa"
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

//处理get请求，da'y
func test()  {
	//
}


