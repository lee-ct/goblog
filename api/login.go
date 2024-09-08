package api

import (
	"go-blog/common"
	"go-blog/service"
	"net/http"
)

func (*ApiHandle) Login(w http.ResponseWriter, req *http.Request) {
	//接收用户名和密码 返回 对应的json数据
	parms := common.GetRequestJsonParam(req)
	userName := parms["username"].(string)
	password := parms["passwd"].(string)
	loginRes, err := service.Login(userName, password)
	if err != nil {
		common.Error(w, err)
	}
	common.Success(w, loginRes)
}
