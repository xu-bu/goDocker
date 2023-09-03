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

var API_URL string
var contractAddress string
var client *ethclient.Client
var auth *bind.TransactOpts
var conn *api.Api

// init是built-in 函数
func initVar() {
	// 获取.env
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error Getwd in main of interactTest/test/test.go:", err)
	}

	envPath := wd + "/../.env"
	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error getting .env in main of interactTest/test/test.go:", err)
	}

	// 注意给全局变量赋值不能用:=
	API_URL, contractAddress = os.Getenv("sepoliaURL"), os.Getenv("CONTRACT_ADDRESS")
	if API_URL == "" || contractAddress == ""  {
		log.Fatal("Error parsing .env in main of interactTest/test/test.go:")
	}

	// 连接到以太坊节点
	client, err = ethclient.Dial(API_URL)
	if err != nil {
		log.Fatal(err, "connection error in deployContract of test.go")
	}

	conn, err = api.NewApi(common.HexToAddress(contractAddress), client)
	if err != nil {
		log.Fatal("new api error in createNFT of test.go\n", err)
	}
}

// 通过私钥算出用户的地址
func getAccountAddress(privateKey string)(string){
	ECDSAPrivateKey, err := crypto.HexToECDSA(privateKey)
    if err != nil {
        log.Fatal("failed to parse private key in getAccountAddress")
    }

	publicKey := ECDSAPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	accountAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	return accountAddress.Hex()
}

func generateAuth(privateKey string)(*bind.TransactOpts){
	initVar()
	ECDSAPrivateKey, err := crypto.HexToECDSA(privateKey)

    if err != nil {
        log.Fatal("failed to parse private key")
    }

    chainID, err := client.ChainID(context.Background())
    if err != nil {
        log.Fatal("failed to get chain ID")
    }

    auth, err := bind.NewKeyedTransactorWithChainID(ECDSAPrivateKey, chainID)
    if err != nil {
        log.Fatal("failed to create transaction options")
    }
	accountAddress:=getAccountAddress(privateKey)
    auth.From = common.HexToAddress(accountAddress)
	return auth
}


// 接口说明
// 用户给出URI和自己的私钥，平台为用户发放对应的NFT
// 前端发起post请求，给出用户私钥privateKey参数和tokenURI
func createNFT(c *gin.Context) {
	initVar()
	
	privateKey,StatusOK := c.GetPostForm("privateKey")
	if !StatusOK{
		panic("未获取到privateKey参数 in createNFT")
	}
	tokenURI,StatusOK := c.GetPostForm("tokenURI")
	if !StatusOK{
		panic("未获取到tokenURI参数 in createNFT")
	}

	auth=generateAuth(privateKey)

	log.Println("create auth successfully")

	_, err := conn.CreateNFT(auth, tokenURI)
	if err != nil {
		log.Fatal("createNFT error in createNFT of test.go\n", err)
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"success",
	})
}


// 接口说明
// 用户将自己的NFT移交给contract并获取ETH
// 前端发起post请求，给出用户私钥privateKey参数 ，想要移交给平台的NFT的tokenID以及price
func mintToContract(c *gin.Context) {
	initVar()
	privateKey,StatusOK := c.GetPostForm("privateKey")
	if !StatusOK{
		panic("未获取到privateKey参数 in mintToContract")
	}
	tokenID,StatusOK := c.GetPostForm("tokenID")
	if !StatusOK{
		panic("未获取到tokenID参数 in mintToContract")
	}
	price,StatusOK := c.GetPostForm("price")
	if !StatusOK{
		panic("未获取到price参数 in mintToContract")
	}
	auth=generateAuth(privateKey)

	// 由于contract的参数是uint256，所以要转一下
	bigIntTokenID := new(big.Int)
	bigIntTokenID.SetString(tokenID, 10) // Base 10
	bigIntPrice := new(big.Int)
	bigIntPrice.SetString(price, 10) // Base 10

	_, err := conn.ListNFT(auth, bigIntTokenID,bigIntPrice) // conn call the balance function of deployed smart contract
	if err != nil {
		log.Fatal("error in mintToContract\n", err)
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"success",
	})
}


// 接口说明
// 用户将自己的NFT移交给contract并获取ETH
// 前端发起post请求，给出用户私钥privateKey和tokenID，返回拥有该token的用户的用户地址
func getOwnerOfNFT(c *gin.Context) {
	privateKey,StatusOK := c.GetPostForm("privateKey")
	if !StatusOK{
		panic("未获取到privateKey参数 in getOwnerOfNFT")
	}
	tokenID,StatusOK := c.GetPostForm("tokenID")
	if !StatusOK{
		panic("未获取到tokenID参数 in getOwnerOfNFT")
	}
	initVar()
	accountAddress:=getAccountAddress(privateKey)
	//调用ownOf的时候不是用auth
	opts := &bind.CallOpts{
		From:    common.HexToAddress(accountAddress), // 发送者地址
	}
	bigIntTokenID := new(big.Int)
	bigIntTokenID.SetString(tokenID, 10) // Base 10
	res, err := conn.OwnerOf(opts, bigIntTokenID) // conn call the balance function of deployed smart contract
	if err != nil {
		log.Println("error in getOwnerOfNFT\n", err)
	}
	c.JSON(http.StatusOK,gin.H{
		"message":res,
	})
}

// 接口说明
// 用户消耗ETH claim NFT
// 前端发起post请求，给出用户私钥privateKey和tokenID
func test(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{
		"message":"success",
	})
}



