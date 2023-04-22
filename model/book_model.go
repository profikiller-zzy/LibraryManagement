package model

import "time"

type BookModel struct {
	MODEL
	BookName        string    `json:"book_name"`        // 图书名
	ISBN            string    `json:"isbn"`             // 图书的ISBN编号
	Author          string    `json:"author"`           // 作者
	Press           string    `json:"press"`            // 出版社
	PublicationDate time.Time `json:"publication_date"` // 出版日期
	Price           float32   `json:"price"`            // 单价
}
