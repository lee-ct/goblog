package router

import (
	"go-blog/api"
	"go-blog/views"
	"net/http"
)

func Router() {
	//登陆
	http.HandleFunc("/login", views.HTML.Login)
	//文章分类
	//http://127.0.0.1:8080/c/2  取后面 1,2
	http.HandleFunc("/c/", views.HTML.Category)
	//1. 页面  views 2. api 数据（json） 3. 静态资源
	http.HandleFunc("/", views.HTML.Index)
	http.HandleFunc("/api", api.Api.SaveAndUpdate)

	http.HandleFunc("/api/v1/login", api.Api.Login)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
