package api_borrow

import (
	"LibraryManagement/global"
	"LibraryManagement/model"
	"LibraryManagement/model/custom_type"
	"LibraryManagement/model/response"
	"LibraryManagement/utils/jwt_util"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// BookRenewView 图书续借
func (BorrowApi) BookRenewView(c *gin.Context) {
	_bookID := c.Param("book_id")
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt_util.CustomUserClaims)

	// 先去查指定的书存不存在
	var bookModel model.BookModel
	err := global.Db.First(&bookModel, _bookID).Error
	if err != nil { // 没有查到
		response.FailWithMessage("该书籍不存在", c)
		return
	}

	// 查询该书是否被借阅
	if bookModel.Status == custom_type.Status(1) {
		response.FailWithMessage("该书籍已经被归还，不可续借", c)
		return
	}

	bookID, err := strconv.ParseUint(_bookID, 10, 32)
	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage("请输入正确的书籍ID", c)
		return
	}

	// 开始事务
	// 更新借阅信息
	// TODO 加上一个判断，判断用户还书时候是否超时
	tx := global.Db.Begin()
	var borrowRecord model.UserBorrowBook
	err = tx.Where("user_id = ? and book_id = ? and return_at is null or return_at = ?", claims.UserID, bookID, sql.NullTime{}).First(&borrowRecord).Error
	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		tx.Rollback() // 回滚事务
		return
	}
	expireAt := borrowRecord.ExpireAt.Add(7 * 24 * time.Hour)
	err = tx.Model(&borrowRecord).Update("expire_at", expireAt).Error
	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		tx.Rollback() // 回滚事务
		return
	}

	// 提交事务
	tx.Commit()

	// 返回成功响应
	response.OKWithMessage(fmt.Sprintf("《%s》续借成功，到期时间延后7天", bookModel.BookName), c)
}
