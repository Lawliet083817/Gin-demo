package model

import "github.com/dgrijalva/jwt-go"

//实现了表单验证功能 binding：“required”

type User struct {
	Username       string `form:"username" json:"username" binding:"required"`
	Password       string `form:"password" json:"password" binding:"required"`
	IdentityNumber string `form:"IdentityNumber" json:"IdentityNumber"`
}

// MyClaims 自定义的JWT声明
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
type Rep struct {
	OldPassword string `form:"old_password" json:"old_password" binding:"required"`
	NewPassword string `form:"old_password" json:"new_password" binding:"required"`
}

type Finder struct {
	Username       string `form:"username" json:"username" binding:"required"`
	IdentityNumber string `form:"IdentityNumber" json:"IdentityNumber" binding:"required"`
}
