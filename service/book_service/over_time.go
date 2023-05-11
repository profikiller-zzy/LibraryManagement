package book_service

import (
	"math"
	"time"
)

// ReturnTimeOut 根据归还时间和借阅期限，计算出超时时间(单位：天)
func ReturnTimeOut(expireAt, returnAt time.Time) int {
	daysLate := int(math.Floor(returnAt.Sub(expireAt).Hours() / 24))
	if daysLate < 0 {
		return 0
	}
	return daysLate
}
