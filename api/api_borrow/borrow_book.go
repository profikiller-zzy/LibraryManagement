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

func (BorrowApi) BookBorrowView(c *gin.Context) {
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
	if bookModel.Status == custom_type.Status(2) {
		response.FailWithMessage("该书籍已经被借阅，请重新选择", c)
		return
	}

	// 将借阅信息写入数据库
	bookID, err := strconv.ParseUint(_bookID, 10, 32)
	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage("请输入正确的书籍ID", c)
		return
	}
	// 开始事务
	tx := global.Db.Begin()
	err = tx.Create(&model.UserBorrowBook{
		UserID:    claims.UserID,
		BookID:    uint(bookID),
		CreatedAt: time.Now(),
		ReturnAt:  sql.NullTime{},
		ExpireAt:  time.Now().Add(7 * 24 * time.Hour), // 默认到期时间为7天
	}).Error
	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		tx.Rollback() // 回滚事务
		return
	}

	// 更新书籍状态
	err = tx.Model(&bookModel).Update("status", custom_type.OnLoan).Error
	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		tx.Rollback() // 回滚事务
		return
	}

	// 提交事务
	tx.Commit()

	// 返回成功响应
	response.OKWithMessage(fmt.Sprintf("《%s》借阅成功", bookModel.BookName), c)
}
