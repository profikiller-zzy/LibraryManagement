package api_book

import (
	"LibraryManagement/global"
	"LibraryManagement/model"
	"LibraryManagement/model/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

type BookQueryByPressRequest struct {
	Press string `json:"press" binding:"required"`
}

func (BookApi) BookQueryByPress(c *gin.Context) {
	var queryPressReq BookQueryByPressRequest
	err := c.ShouldBindJSON(&queryPressReq)
	if err != nil {
		response.FailBecauseOfParamError(err, &queryPressReq, c)
		return
	}

	var bookList []model.BookModel
	queryName := "%" + queryPressReq.Press + "%"
	tx := global.Db.Where("press like ?", queryName).Find(&bookList)
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
