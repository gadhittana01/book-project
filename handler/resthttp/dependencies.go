package resthttp

import (
	"context"

	"gihub.com/gadhittana01/book-project/services"
)

type (
	BookService interface {
		GetListOfBooks(ctx context.Context, req services.GetListOfBooksReq) (services.GetListOfBooksResp, error)
		BorrowBook(ctx context.Context, req services.BorrowBookReq) (services.BorrowBookRes, error)
		GetBookReservation(ctx context.Context, req services.GetBookReservationReq) (map[int][]services.GetBookReservationRes, error)
	}
)
