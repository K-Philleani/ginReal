package controller

import (
	"fmt"
	"ginReal/common"
	"ginReal/model"
	"ginReal/response"
	"ginReal/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)


func isPhoneExist(db *gorm.DB, phone string) bool{
	var user model.User
	db.Where("phone=?", phone).First(&user)
	fmt.Println("user:", user)
	if user.ID != 0 {
		return true
	}
	return false
}

func Register(ctx *gin.Context) {
	var requestUser = model.User{}
	db := common.DB
	ctx.BindJSON(&requestUser)
	name := requestUser.Name
	phone := requestUser.Phone
	password := requestUser.Password
	if len(phone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 442, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码必须大于6位")
		return
	}
	if len(name) == 0 {
		name = utils.RandomString(10)
	}
	if isPhoneExist(db, phone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已经存在")
		return
	}
	hasePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "加密错误")
		return
	}
	newUser := model.User{
		Name:     name,
		Phone:    phone,
		Password: string(hasePassword),
	}
	db.Create(&newUser)
	// 发送token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "系统异常")
	}
	response.Success(ctx, gin.H{ "token": token }, "注册成功")

}

func Login(ctx *gin.Context) {
	db := common.DB
	var u model.User
	ctx.BindJSON(&u)
	if len(u.Phone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{"phone": u.Phone}, "手机号必须11位")
		return
	}
	if len(u.Password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码必须大于6位")
		return
	}
	var user model.User
	db.Where("phone", u.Phone).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password)); err != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
		return
	}
	response.Success(ctx, gin.H{"token": token}, "登陆成功")
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.Response(ctx, http.StatusOK, 200, gin.H{"user": user}, "用户信息")
}

