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

type BookCreateRequest struct {
	BookName        string `json:"book_name" binding:"required" msg:"请输入正确的图书名"`         // 图书名
	ISBN            string `json:"isbn" binding:"required" msg:"请输入正确的图书的ISBN编号"`        // 图书的ISBN编号
	Author          string `json:"author" binding:"required" msg:"请输入正确的作者"`             // 作者
	Press           string `json:"press" binding:"required" msg:"请输入正确的出版社"`             // 出版社
	PublicationDate string `json:"publication_date" binding:"required" msg:"请输入正确的出版日期"` // 出版日期
	Price           string `json:"price" binding:"required" msg:"请输入正确的单价"`              // 单价
}

func (BookApi) BookCreateView(c *gin.Context) {
	var bcReq BookCreateRequest
	err := c.ShouldBindJSON(&bcReq)
	if err != nil {
		response.FailBecauseOfParamError(err, &bcReq, c)
		return
	}

	publicationDate, _ := time.Parse("2006-01-02", bcReq.PublicationDate)
	priceFloat64, err := strconv.ParseFloat(bcReq.Price, 64)
	if err != nil {
		response.FailWithMessage("无法解析价格，请你输入合法的价格", c)
		return
	}
	err = global.Db.Create(&model.BookModel{
		BookName:        bcReq.BookName,
		ISBN:            bcReq.ISBN,
		Author:          bcReq.Author,
		Press:           bcReq.Press,
		PublicationDate: publicationDate,
		Price:           priceFloat64,
	}).Error
	if err != nil { // 书籍添加失败
		global.Log.Error(err.Error())
		response.FailWithMessage(fmt.Sprintf("添加书籍失败，报错信息:%s", err.Error()), c)
		return
	}
	response.OKWithMessage(fmt.Sprintf("书籍《%s》添加成功", bcReq.BookName), c)
}
