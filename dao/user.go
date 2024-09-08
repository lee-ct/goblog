package dao

import (
	"errors"
	"go-blog/model"
	"go-blog/utils"
	"log"
)

func GetUserNameBYId(userId int) string {
	row := DB.QueryRow("select user_name from blog_user where uid=?", userId)

	var userName string
	row.Scan(&userName)

	return userName
}

func GetUser(username, password string) *model.User {
	password = utils.Md5Crypt(password)
	row := DB.QueryRow("select * from blog_user where user_name=? and passwd=?", username, password)
	var user *model.User
	user = &model.User{}
	err := row.Scan(&user.Uid, &user.UserName, &user.Passwd, &user.Avatar, &user.CreateAt, &user.UpdateAt)
	if err != nil {
		log.Println(errors.New("GetUser错误，查不到此用户"), err)
		return nil
	}
	return user
}
