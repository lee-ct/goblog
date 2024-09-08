package service

import (
	"errors"
	"go-blog/dao"
	"go-blog/model"
	"go-blog/utils"
	"log"
)

func Login(username, password string) (*model.LoginRes, error) {

	user := dao.GetUser(username, password)
	if user == nil {
		return nil, errors.New("账号密码不正确")
	}
	//根据用户id生成token
	token, err := utils.Award(&user.Uid)
	if err != nil {
		log.Println("jwt 未能生成")
	}

	var userInfo model.UserInfo
	userInfo.UserName = user.UserName
	userInfo.Avatar = user.Avatar
	userInfo.Uid = user.Uid

	var ls = &model.LoginRes{
		token,
		userInfo,
	}
	return ls, nil
}
