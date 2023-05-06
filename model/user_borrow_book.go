package model

import (
	"database/sql"
	"time"
)

// UserBorrowBook 读者借阅图书的自定义关联表
type UserBorrowBook struct {
	ID        uint         `gorm:"primaryKey"` // 标记借阅记录的唯一ID
	UserID    uint         `json:"user_id"`    // 借阅图书的用户ID
	User      UserModel    // 借阅图书的用户
	BookID    uint         `json:"book_id"` // 借阅图书ID
	Book      BookModel    // 借阅的图书
	CreatedAt time.Time    `json:"created_at"`                    // 借书时间
	ReturnAt  sql.NullTime `gorm:"default:null" json:"return_at"` // 还书时间，如果未换则为空
}
