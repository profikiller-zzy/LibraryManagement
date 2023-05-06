package model

import (
	"LibraryManagement/model/custom_type"
	"time"
)

type BookModel struct {
	MODEL
	BookName        string             `json:"book_name"`        // 图书名
	ISBN            string             `json:"isbn"`             // 图书的ISBN编号
	Author          string             `json:"author"`           // 作者
	Press           string             `json:"press"`            // 出版社
	PublicationDate time.Time          `json:"publication_date"` // 出版日期
	Price           float64            `json:"price"`            // 单价
	Status          custom_type.Status `json:"status"`           // 书籍的状态： 1 被借阅 2 空闲
}
