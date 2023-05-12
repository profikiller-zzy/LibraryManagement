package router

import (
	"LibraryManagement/api/api_book"
	middleware "LibraryManagement/middleware/jwt"
)

func (r RGroup) BookRouter() {
	var BookApiApp api_book.BookApi
	r.GET("/book_list/", BookApiApp.BookListView)
	r.PUT("/book_create/", middleware.JwtAdmin(), BookApiApp.BookCreateView)
	r.POST("/book_update/:id", middleware.JwtAdmin(), BookApiApp.BookUpdateView)
	r.DELETE("/book_remove/", middleware.JwtAdmin(), BookApiApp.AdRemoveView)
	r.POST("/book_query_book_name/", BookApiApp.BookQueryByBookName)
	r.POST("/book_query_author/", BookApiApp.BookQueryByAuthor)
	r.POST("/book_query_press/", BookApiApp.BookQueryByPress)
}
