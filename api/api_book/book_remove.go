package api_book

import (
	"LibraryManagement/global"
	"LibraryManagement/model"
	"LibraryManagement/model/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (BookApi) AdRemoveView(c *gin.Context) {
	var rmReq model.RemoveRequest
	var bookList []model.BookModel
	var count int64 = 0

	err := c.ShouldBindJSON(&rmReq)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("参数绑定失败,报错信息：%s", err.Error()), c)
		return
	}

	count = global.Db.Find(&bookList, rmReq.IDList).RowsAffected
	if count == 0 { // 需要删除的图片ID没有在数据库中查到
		response.FailWithMessage("文件不存在", c)
		return
	}
	global.Db.Delete(&model.BookModel{}, rmReq.IDList)
	response.FailWithMessage(fmt.Sprintf("删除 %d 本书籍记录成功", count), c)
}
