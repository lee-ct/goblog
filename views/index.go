package views

import (
	"go-blog/common"
	"go-blog/service"
	"log"
	"net/http"
	"strconv"
)

func (*HtmlApi) Index(w http.ResponseWriter, req *http.Request) {
	index := common.Template.Index
	err := req.ParseForm()
	if err != nil {
		log.Println("表单解析错误", err)
	}
	pageStr := req.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	//每页显示的数量
	pageSize := 10
	hr, err := service.GetAllIndexInfo(page, pageSize)
	if err != nil {
		log.Println("GetAllIndexInfo出错了")
	}

	//上面页面涉及的变量下面都要有
	index.WriteData(w, hr)
}
