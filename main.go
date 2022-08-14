package main

import (
	"leadboard/config"
	"leadboard/model"
	"leadboard/route"
)

func main() {
	r := route.InitRoute()
	model.BuildConnection(config.Parse())
	r.Run(":10001") //TODO:运行到容器开放的端口上
}
