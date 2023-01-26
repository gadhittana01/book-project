package services

import (
	"context"

	"gihub.com/gadhittana01/book-project/pkg/domain"
)

type (
	BookResource interface {
		GetListOfBooks(ctx context.Context, req domain.GetListOfBooksReq) (domain.GetListOfBooksResp, error)
		BorrowBook(ctx context.Context, req domain.BorrowBookReq) error
		GetBookByKey(ctx context.Context, req domain.GeBookByKeyReq) (domain.Book, error)
		GetBookReservation(ctx context.Context, req domain.GetBookReservationReq) (map[int][]domain.BorrowBookReq, error)
	}
)
