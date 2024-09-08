package views

import (
	"fmt"
	"go-blog/common"
	"go-blog/service"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (*HtmlApi) Category(w http.ResponseWriter, req *http.Request) {
	categoryTemplate := common.Template.Category
	path := req.URL.Path
	cIdstr := strings.TrimPrefix(path, "/c/")
	cId, err := strconv.Atoi(cIdstr)
	if err != nil {
		log.Println("不能识别此路径", err)
	}

	pageStr := req.Form.Get("page")
	fmt.Println(pageStr)
	if pageStr == "" {
		pageStr = "1"
	}

	page, _ := strconv.Atoi(pageStr)
	//显示每页数量
	pageSize := 10

	CategoryResponse, err := service.GetPostsByCategoryId(cId, page, pageSize)
	if err != nil {
		log.Println("获取分类表单失败GetPostsByCategoryId", err)
		return
	}

	categoryTemplate.WriteData(w, CategoryResponse)
}
