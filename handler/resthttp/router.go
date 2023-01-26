package resthttp

import (
	"github.com/go-chi/chi"
)

type RouterDependencies struct {
	BS BookService
}

func NewRoutes(rd RouterDependencies) *chi.Mux {
	router := chi.NewRouter()

	bh := newBookHandler(rd.BS)
	router.Get("/get-books", bh.GetListOfBooks)
	router.Post("/borrow-book", bh.BorrowBook)
	router.Get("/get-book-reservation", bh.GetBookReservation)

	return router
}
