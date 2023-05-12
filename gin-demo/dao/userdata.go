package dao

import (
	"awesomeProject/gin-demo/utils"
	"github.com/gin-gonic/gin"
)

var database1 = map[string]string{
	"WJQ": "118817",
}

var database2 = map[string]string{
	"WJQ": "083817",
}

// AddUser 添加用户
func AddUser(username, password, IdentityNumber string) {
	database1[username] = password
	database2[username] = IdentityNumber
}

// SelectUser 检查是否存在该用户
func SelectUser(username string) bool {
	if database1[username] == "" {
		return false
	} else {
		return true
	}
}

// GetPassword 获取用户的密码（服务端）
func GetPassword(username string) string {
	return database1[username]
}

// ChangePassword 修改密码
func ChangePassword(username string, NewPassword string) {
	database1[username] = NewPassword
	return
}

// FindPassword 查找密码
func FindPassword(c *gin.Context, username, IdentityNumber string) string {
	//查看该用户是否存在
	if SelectUser(username) == false {
		utils.FailRes(c, "this user doesn't exist")
		return "doesn't exist"
	}
	//检验身份是否正确
	if IdentityNumber == database2[username] {
		return GetPassword(username)
	} else {
		utils.FailRes(c, "IdentityNumber isn't match with our database")
		return "wrong number"
	}

}
