package main

import (
	"go-blog/common"
	"go-blog/router"
	"log"
	"net/http"
)

// 显示模版
func init() {
	//模版加载
	common.LoadTemplate()
}

func main() {
	serve := http.Server{
		Addr: "127.0.0.1:8080",
	}
	router.Router()
	err := serve.ListenAndServe()
	if err != nil {
		log.Fatalf("ListenAndServe:" + err.Error())
	}
}
