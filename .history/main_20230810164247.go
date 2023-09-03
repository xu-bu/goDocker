package main

import (
	_ "context"
	_ "errors"
	_ "log"
	_ "net/http"
	_ "os"
	_ "os/signal"
	_ "syscall"
	_ "time"

	"github.com/gin-gonic/gin"
	_ "github.com/julienschmidt/httprouter"
)

func main() {
	// 创建一个默认的路由引擎
	app := gin.Default()

	// post 请求在 runApi中要使用body中的x-www-form-urlencoded类型的数据
	app.Handle("POST", "/", createNFT)
	// app.GET("/", test)

	// 接口说明
	// 用户将自己的NFT移交给contract并获取ETH
	// 前端发起post请求，给出用户私钥privateKey参数 ，想要移交给平台的NFT的tokenID以及price
	app.Handle("POST", "/mintContract", mintToContract)

	// 接口说明
	// 获取对应tokenID的拥有者的address
	// 前端发起post请求，给出用户私钥privateKey和tokenID，返回拥有该token的用户的用户地址
	app.Handle("POST", "/ownerOf", getOwnerOfNFT)

	// 接口说明
	// 用户消耗ETH claim NFT
	// 前端发起post请求，给出用户私钥privateKey和tokenID
	app.Handle("POST", "/claim", claimNFT)
	// 默认8080端口
	app.Run()
}
