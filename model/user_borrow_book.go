package model

import "time"

// UserBorrowBook 读者借阅图书的自定义关联表
type UserBorrowBook struct {
	UserID    uint      `gorm:"primaryKey"`        // 借阅图书的用户ID
	User      UserModel `gorm:"foreignKey:UserID"` // 借阅图书的用户
	BookID    uint      `gorm:"primaryKey"`        // 借阅图书ID
	Book      BookModel `gorm:"foreignKey:BookID"` // 借阅的图书
	CreatedAt time.Time `json:"created_at"`        // 创建时间
}
