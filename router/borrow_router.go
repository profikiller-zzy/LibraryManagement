package router

import (
	"LibraryManagement/api/api_borrow"
	middleware "LibraryManagement/middleware/jwt"
)

func (r RGroup) BorrowRouter() {
	var BorrowApiApp api_borrow.BorrowApi
	r.PUT("/book_borrow/:book_id", middleware.JwtAuth(), BorrowApiApp.BookBorrowView)
	r.POST("/book_return/:book_id", middleware.JwtAuth(), BorrowApiApp.BookReturnView)
	r.POST("/book_renew/:book_id", middleware.JwtAuth(), BorrowApiApp.BookRenewView)
}
