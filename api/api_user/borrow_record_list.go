package api_user

import (
	"LibraryManagement/global"
	"LibraryManagement/model"
	"LibraryManagement/model/response"
	"LibraryManagement/utils/jwt_util"
	"database/sql"
	"github.com/gin-gonic/gin"
	"time"
)

type BorrowRecordResponse struct {
	BookID    uint            `json:"book_id"` // 借阅图书ID
	Book      model.BookModel // 借阅的图书
	CreatedAt time.Time       `json:"created_at"`                    // 借书时间
	ReturnAt  sql.NullTime    `gorm:"default:null" json:"return_at"` // 还书时间，如果未换则为空
}

// BorrowRecordList 返回用户的所有借阅记录
func (UserApi) BorrowRecordList(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt_util.CustomUserClaims)

	var borrowRecordList []BorrowRecordResponse
	// 用到了GORM的智能选择字段
	err := global.Db.Model(&model.UserBorrowBook{}).Where("user_id = ? order by created_at desc", claims.UserID).Preload("Book").Take(&borrowRecordList).Error
	if err != nil {
		response.FailWithMessage("你还没有借过书呢，快去借阅书籍吧", c)
		return
	}
	response.OKWithData(borrowRecordList, c)
}
