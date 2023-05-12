package api

import (
	"awesomeProject/gin-demo/api/middleware"
	"awesomeProject/gin-demo/dao"
	"awesomeProject/gin-demo/model"
	"awesomeProject/gin-demo/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

// 登录接口的实现：
func login(c *gin.Context) {

	var u model.User
	if err := c.ShouldBind(&u); err != nil {
		utils.FailRes(c, "verification failed")
		return
	} //表单验证失败的错误处理

	//检验账户是否存在
	Exist := dao.SelectUser(u.Username)

	//错误处理：
	//账户不存在
	if !Exist {
		utils.FailRes(c, "this user doesn't exist!")
		return
	}
	TurePassword := dao.GetPassword(u.Username)
	//密码错误
	if u.Password != TurePassword {
		utils.FailRes(c, "this password is wrong!")
		return
	}

	//登录成功，需要JWT认证
	Claims := model.MyClaims{
		Username: u.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "WJQ",
		},
	}

	//使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodES256, Claims)
	//使用指定secret进行签名并获得编码后的token
	tokenString, _ := token.SignedString(middleware.Secret)
	utils.SucRes(c, tokenString)
}

// 注册接口的实现：
func register(c *gin.Context) {
	//ShouldBind 绑定表单 错误处理
	var u model.User
	if err := c.ShouldBind(&u); err != nil {
		utils.SucRes(c, "verification failed")
		return
	}

	//验证用户名是否已经存在
	Exist := dao.SelectUser(u.Username)
	if Exist {
		utils.FailRes(c, "user has already existed")
		return
	}

	//调用数据库dao中的方法
	dao.AddUser(u.Username, u.Password, u.IdentityNumber)

	//返回JSON格式信息
	utils.SucRes(c, "add user successful")
}

// 从token获得username
func getUsernameFromToken(c *gin.Context) {
	username, _ := c.Get("username")
	utils.SucRes(c, username.(string))
}

// 修改密码
func resetPassword(c *gin.Context) {
	username1, _ := c.Get("username")
	username := username1.(string) //类型断言将any修改为string便于后续使用

	var ret model.Rep
	if err := c.ShouldBind(&ret); err != nil {
		utils.FailRes(c, "binding error!")
		return
	}

	//老密码输入错误或者与新密码相同
	if ret.OldPassword != dao.GetPassword(username) {
		utils.FailRes(c, "wrong old password!")
		return
	}
	if ret.NewPassword == ret.OldPassword {
		utils.FailRes(c, "your new password is same as your old one!")
		return
	}
	//调用dao中修改函数
	dao.ChangePassword(username, ret.NewPassword)
	utils.SucRes(c, "reset is successful")
}

func findPassword(c *gin.Context) {

	var fd model.Finder
	if err := c.ShouldBind(&fd); err != nil {
		utils.FailRes(c, "binding error!")
		return
	}

	upassword := dao.FindPassword(c, fd.Username, fd.IdentityNumber)
	utils.SucRes(c, upassword)
	return
}
