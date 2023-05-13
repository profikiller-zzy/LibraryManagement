package api_book

import (
	"LibraryManagement/global"
	"LibraryManagement/model"
	"LibraryManagement/model/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

type BookQueryByAuthorRequest struct {
	Author string `json:"author" binding:"required"`
}

func (BookApi) BookQueryByAuthor(c *gin.Context) {
	var queryAuthorReq BookQueryByAuthorRequest
	err := c.ShouldBindJSON(&queryAuthorReq)
	if err != nil {
		response.FailBecauseOfParamError(err, &queryAuthorReq, c)
		return
	}

	var bookList []model.BookModel
	queryName := "%" + queryAuthorReq.Author + "%"
	tx := global.Db.Where("author like ?", queryName).Find(&bookList)
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
	response.OKWithPagingData(bookList, count, c)
	return
}
