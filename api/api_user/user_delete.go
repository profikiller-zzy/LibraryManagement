package api_user

import (
	"LibraryManagement/global"
	"LibraryManagement/model"
	"LibraryManagement/model/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (UserApi) DeleteUserView(c *gin.Context) {
	userID := c.Param("id")

	var userModel model.UserModel
	err := global.Db.First(&userModel, userID).Error
	if err != nil { // 没有找到符合条件的记录
		response.FailWithMessage("该用户不存在", c)
		return
	}

	var borrowRecord model.UserBorrowBook
	err = global.Db.First(&borrowRecord, "user_id = ? and return_at is null", userID).Error
	if err == nil { // 找到符合条件的记录
		response.FailWithMessage(fmt.Sprintf("该用户借阅图书 %d ，未归还，不可删除", borrowRecord.BookID), c)
		return
	}

	err = global.Db.Delete(&userModel).Error
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，报错信息: %s", err.Error()), c)
		return
	}
	response.OKWithMessage(fmt.Sprintf("删除用户 %d 成功", userModel.ID), c)
}
