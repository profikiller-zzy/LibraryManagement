package api_book

import (
	"LibraryManagement/global"
	"LibraryManagement/model"
	"LibraryManagement/model/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type BookUpdateRequest struct {
	BookName        string `json:"book_name" binding:"required" msg:"请输入正确的图书名"`         // 图书名
	ISBN            string `json:"isbn" binding:"required" msg:"请输入正确的图书的ISBN编号"`        // 图书的ISBN编号
	Author          string `json:"author" binding:"required" msg:"请输入正确的作者"`             // 作者
	Press           string `json:"press" binding:"required" msg:"请输入正确的出版社"`             // 出版社
	PublicationDate string `json:"publication_date" binding:"required" msg:"请输入正确的出版日期"` // 出版日期
	Price           string `json:"price" binding:"required" msg:"请输入正确的单价"`              // 单价
}

func (BookApi) BookUpdateView(c *gin.Context) {
	var buReq BookUpdateRequest
	err := c.ShouldBindJSON(&buReq)
	if err != nil {
		response.FailBecauseOfParamError(err, &buReq, c)
		return
	}

	bookID := c.Param("id")

	var bookModel model.BookModel
	err = global.Db.First(&bookModel, "id = ?", bookID).Error
	if err != nil { // 没有找到符合条件的记录
		response.FailWithMessage("该书籍不存在", c)
		return
	}

	publicationDate, _ := time.Parse("2006-01-02", buReq.PublicationDate)
	priceFloat64, err := strconv.ParseFloat(buReq.Price, 64)
	err = global.Db.Model(&bookModel).Updates(map[string]interface{}{
		"book_name":        buReq.BookName,
		"isbn":             buReq.ISBN,
		"author":           buReq.Author,
		"press":            buReq.Press,
		"publication_date": publicationDate,
		"price":            priceFloat64,
	}).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithMessage(fmt.Sprintf("修改书籍《%s》成功", buReq.BookName), c)
}
