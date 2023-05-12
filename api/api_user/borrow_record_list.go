package api_user

import (
	"LibraryManagement/global"
	"LibraryManagement/model"
	"LibraryManagement/model/custom_type"
	"LibraryManagement/model/response"
	"LibraryManagement/utils/jwt_util"
	"database/sql"
	"github.com/gin-gonic/gin"
	"time"
)

type BorrowRecord struct {
	BookID    uint            `json:"book_id"` // 借阅图书ID
	Book      model.BookModel // 借阅的图书
	CreatedAt time.Time       `json:"created_at"` // 借书时间
	ReturnAt  sql.NullTime    `json:"return_at"`  // 还书时间，如果未换则为空
	ExpireAt  time.Time       `json:"expire_at"`  // 到期时间
}

type BorrowRecordResponse struct {
	ID              uint               `json:"id"`                            // 图书编号
	BookName        string             `json:"book_name"`                     // 图书名
	ISBN            string             `json:"isbn"`                          // 图书的ISBN编号
	Author          string             `json:"author"`                        // 作者
	Press           string             `json:"press"`                         // 出版社
	PublicationDate time.Time          `json:"publication_date"`              // 出版日期
	Price           float64            `json:"price"`                         // 单价
	Status          custom_type.Status `json:"status"`                        // 书籍的状态： 1 被借阅 2 空闲
	CreatedAt       time.Time          `json:"created_at"`                    // 借书时间
	ReturnAt        sql.NullTime       `gorm:"default:null" json:"return_at"` // 还书时间，如果未换则为空
	ExpireAt        time.Time          `json:"expire_at"`                     // 到期时间
}

// BorrowRecordList 返回用户的所有借阅记录
func (UserApi) BorrowRecordList(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt_util.CustomUserClaims)

	var borrowRecordList []BorrowRecord
	// 用到了GORM的智能选择字段
	tx := global.Db.Debug().Model(&model.UserBorrowBook{}).Where("user_id = ? order by created_at desc", claims.UserID).Preload("Book").Find(&borrowRecordList)
	err := tx.Error
	count := tx.RowsAffected
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if count == 0 {
		response.FailWithMessage("你还没有借过书呢，快去借阅书籍", c)
		return
	}

	var res = make([]BorrowRecordResponse, count)
	for index, value := range borrowRecordList {
		res[index] = BorrowRecordResponse{
			ID:              value.Book.ID,
			BookName:        value.Book.BookName,
			ISBN:            value.Book.ISBN,
			Author:          value.Book.Author,
			Press:           value.Book.Press,
			PublicationDate: value.Book.PublicationDate,
			Price:           value.Book.Price,
			Status:          value.Book.Status,
			CreatedAt:       value.CreatedAt,
			ReturnAt:        value.ReturnAt,
			ExpireAt:        value.ExpireAt,
		}
	}
	response.OKWithPagingData(res, count, c)
}
