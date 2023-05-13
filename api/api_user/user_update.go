package api_user

import (
	"LibraryManagement/global"
	"LibraryManagement/model"
	"LibraryManagement/model/custom_type"
	"LibraryManagement/model/response"
	"LibraryManagement/utils/jwt_util"
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

type UserUpdateRequest struct {
	NickName        string `json:"nick_name"`        // 昵称
	TelephoneNumber string `json:"telephone_number"` // 读者的电话号码
	Gender          string `json:"gender"`           // 性别
	Age             int    `json:"age"`              // 读者年龄
}

type UserUpdateInfo struct {
	NickName        string             `json:"nick_name"`        // 昵称
	TelephoneNumber string             `json:"telephone_number"` // 读者的电话号码
	Gender          custom_type.Gender `json:"gender"`           // 性别
	Age             int                `json:"age"`              // 读者年龄
}

func (UserApi) UserUpdateView(c *gin.Context) {
	var upReq UserUpdateRequest
	err := c.ShouldBindJSON(&upReq)
	// 判断参数是否合法
	if err != nil {
		response.FailBecauseOfParamError(err, &upReq, c)
		return
	}

	// 获取用户的ID和昵称
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt_util.CustomUserClaims)

	var userModel model.UserModel
	err = global.Db.First(&userModel, claims.UserID).Error
	if err != nil {
		response.FailWithMessage("用户不存在", c)
		return
	}

	upInfo := UserUpdateInfo{
		NickName:        upReq.NickName,
		TelephoneNumber: upReq.TelephoneNumber,
		Age:             upReq.Age,
	}
	switch upReq.Gender {
	case "男性":
		upInfo.Gender = custom_type.Male
	case "女性":
		upInfo.Gender = custom_type.Female
	default:
		upInfo.Gender = custom_type.Male
	}

	upReqMap := structs.Map(&upInfo)
	err = global.Db.Model(&userModel).Updates(upReqMap).Error
	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage(fmt.Sprintf("用户消息修改失败,%s", err.Error()), c)
		return
	}
	response.OKWithMessage("修改个人消息成功", c)
}
