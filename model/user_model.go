package model

import "LibraryManagement/model/custom_type"

type UserModel struct {
	MODEL
	UserName        string             `gorm:"size:36" json:"user_name"`   // 用户名
	Password        string             `gorm:"size:36" json:"password"`    // 密码
	NickName        string             `gorm:"size:36" json:"nick_name"`   // 昵称
	TelephoneNumber string             `gorm:"16" json:"telephone_number"` // 读者的电话号码
	Gender          custom_type.Gender `gorm:"size:4" json:"gender"`       // 性别
	Age             int                `gorm:"size:4" json:"age"`          // 读者年龄
	BorrowedBooks   []BookModel        `gorm:"many2many:user_borrow_books;joinForeignKey:UserID;joinReferences:BookID" json:"borrowed_books"`
}
