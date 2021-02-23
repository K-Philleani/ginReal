package controller

import (
	"ginReal/common"
	"ginReal/model"
	"ginReal/response"
	"github.com/gin-gonic/gin"
)

func GetUserList(ctx *gin.Context) {
	var userList []model.User
	result := common.DB.Find(&userList)
	var list []map[string]interface{}
	for _, v := range userList {
		temp := map[string]interface{} {
			"id": v.ID,
			"name": v.Name,
			"phone": v.Phone,
		}
		list = append(list, temp)
	}
	response.Success(ctx,
		gin.H{
			"count": result.RowsAffected,
			"userList": list,
		},
	"success",
	)
}

func DeleteUserByPhone(ctx *gin.Context) {
	var user model.User
	ctx.BindJSON(&user)
	if user.Phone == "" {
		response.Fail(ctx, nil, "未获取phone")
		return
	}
	if !isPhoneExist(common.DB, user.Phone) {
		response.Fail(ctx, nil, "phone不存在")
		return
	}
	currentPhonne, _ := ctx.Get("currentPhonne")
	if currentPhonne == user.Phone {
		response.Fail(ctx, nil, "禁止删除当前登录用户账号")
		return
	}
	res := common.DB.Where("phone=?", user.Phone).Delete(&model.User{})
	if res.Error != nil {
		response.Fail(ctx, nil, "删除失败")
		return
	}
	response.Success(ctx, gin.H{"num": res.RowsAffected}, "删除成功")
}