package api_book

import (
	"LibraryManagement/global"
	"LibraryManagement/model"
	"LibraryManagement/model/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

type BookQueryByBookNameRequest struct {
	BookName string `json:"book_name" binding:"required"`
}

func (BookApi) BookQueryByBookName(c *gin.Context) {
	var queryNameReq BookQueryByBookNameRequest
	err := c.ShouldBindJSON(&queryNameReq)
	if err != nil {
		response.FailBecauseOfParamError(err, &queryNameReq, c)
		return
	}

	var bookList []model.BookModel
	queryName := "%" + queryNameReq.BookName + "%"
	tx := global.Db.Where("book_name like ?", queryName).Find(&bookList)
	err = tx.Error
	count := tx.RowsAffected
	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage(fmt.Sprintf("查询出错!报错信息：%s", err.Error()), c)
		return
	}
	if count == 0 {
		response.OKWithMessage("没有相关的书籍", c)
		return
	}
	response.OKWithData(bookList, c)
	return
}
