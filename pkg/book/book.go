package book

import (
	"context"

	"gihub.com/gadhittana01/book-project/config"
	"gihub.com/gadhittana01/book-project/pkg/domain"
)

type IResource interface {
	GetListOfBooks(ctx context.Context, req domain.GetListOfBooksReq) (domain.GetListOfBooksResp, error)
	BorrowBook(ctx context.Context, req domain.BorrowBookReq) error
	GetBookByKey(ctx context.Context, req domain.GeBookByKeyReq) (domain.Book, error)
	GetBookReservation(ctx context.Context, req domain.GetBookReservationReq) (map[int][]domain.BorrowBookReq, error)
}

type module struct {
	external   external
	persistent persistent
}

func New(cfg *config.GlobalConfig, httpclient HttpResource) IResource {
	return &module{
		external:   newExternal(cfg, httpclient),
		persistent: newPersistent(),
	}
}

func (m module) GetListOfBooks(ctx context.Context, req domain.GetListOfBooksReq) (domain.GetListOfBooksResp, error) {
	return m.external.getListOfBooks(ctx, req)
}

func (m module) BorrowBook(ctx context.Context, req domain.BorrowBookReq) error {
	return m.persistent.borrowBook(ctx, req)
}

func (m module) GetBookByKey(ctx context.Context, req domain.GeBookByKeyReq) (domain.Book, error) {
	return m.external.getBookByKey(ctx, req)
}

func (m module) GetBookReservation(ctx context.Context, req domain.GetBookReservationReq) (map[int][]domain.BorrowBookReq, error) {
	return m.persistent.getBookReservation(ctx, req)
}
