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
	// app.GET("/", test)
	// 默认8080端口
	app.Run()
}
