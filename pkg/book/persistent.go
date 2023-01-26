package book

import (
	"context"
	"errors"

	"gihub.com/gadhittana01/book-project/pkg/domain"
)

type persistent interface {
	borrowBook(ctx context.Context, req domain.BorrowBookReq) error
	getBookReservation(ctx context.Context, req domain.GetBookReservationReq) (map[int][]domain.BorrowBookReq, error)
}

type persistentModule struct {
}

func newPersistent() persistent {
	return &persistentModule{}
}

var books map[int][]domain.BorrowBookReq = make(map[int][]domain.BorrowBookReq)

func (m *persistentModule) borrowBook(ctx context.Context, req domain.BorrowBookReq) error {
	if req.UserID == 0 {
		return errors.New("User ID is empty")
	}
	books[req.UserID] = append(books[req.UserID], req)
	return nil
}

func (m *persistentModule) getBookReservation(ctx context.Context, req domain.GetBookReservationReq) (map[int][]domain.BorrowBookReq, error) {
	if req.UserID == 0 {
		return books, nil
	}
	return map[int][]domain.BorrowBookReq{
		req.UserID: books[req.UserID],
	}, nil
}
