package api_user

import (
	"LibraryManagement/global"
	"LibraryManagement/model"
	"LibraryManagement/model/custom_type"
	"LibraryManagement/model/response"
	"LibraryManagement/utils/jwt_util"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserInfoResponse struct {
	NickName        string             `json:"nick_name"`        // 昵称
	TelephoneNumber string             `json:"telephone_number"` // 读者的电话号码
	Gender          custom_type.Gender `json:"gender"`           // 性别
	Age             int                `json:"age"`              // 读者年龄
}

func (UserApi) UserInfoView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt_util.CustomUserClaims)

	var userModel model.UserModel
	err := global.Db.First(&userModel, claims.UserID).Error
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败, 报错信息：%s", err.Error()), c)
		return
	}
	var UserInfoRep = UserInfoResponse{
		NickName:        userModel.NickName,
		TelephoneNumber: userModel.TelephoneNumber,
		Gender:          userModel.Gender,
		Age:             userModel.Age,
	}
	response.OKWithData(UserInfoRep, c)
}
