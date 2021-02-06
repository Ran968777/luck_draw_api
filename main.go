package main

import (
	"fmt"
	"luck_draw/routers"
	"luck_draw/service"
	"luck_draw/util"
	"net/http"
)

func main() {
	router:= routers.InitRouter()

	service.Pay()

	//go queue.Listen()
	//go queue.ScanActivity()

	server := &http.Server{
		Addr:           ":8081",
		Handler:        router,
	}

	err := server.ListenAndServe()
	if err != nil {
		util.Error(fmt.Sprintf("启动服务失败：%v\n",err))
	}
}
