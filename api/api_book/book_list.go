package api_book

import (
	"LibraryManagement/global"
	"LibraryManagement/model"
	"LibraryManagement/model/response"
	"LibraryManagement/service/common_service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (BookApi) BookListView(c *gin.Context) {
	var pageModel model.PageInfo
	err := c.ShouldBindQuery(&pageModel)
	if err != nil {
		response.FailWithCode(response.ParameterError, c)
		return
	}

	var bookList []model.BookModel
	var count int64
	bookList, count, err = common_service.PagingList(model.BookModel{}, common_service.PageInfoDebug{
		PageInfo: pageModel,
		Debug:    true,
	})

	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage(fmt.Sprintf("查询失败，报错信息:%s", err.Error()), c)
		return
	}
	response.OKWithPagingData(bookList, count, c)
	return
}
