package services

import (
	"context"

	"gihub.com/gadhittana01/book-project/pkg/domain"
)

type BookService interface {
	GetListOfBooks(ctx context.Context, req GetListOfBooksReq) (GetListOfBooksResp, error)
	BorrowBook(ctx context.Context, req BorrowBookReq) (BorrowBookRes, error)
	GetBookReservation(ctx context.Context, req GetBookReservationReq) (map[int][]GetBookReservationRes, error)
}

type bookService struct {
	br BookResource
}

func NewBookService(dep BookDependencies) (BookService, error) {
	return &bookService{
		br: dep.BR,
	}, nil
}

func (p bookService) GetListOfBooks(ctx context.Context, req GetListOfBooksReq) (GetListOfBooksResp, error) {
	var result GetListOfBooksResp

	res, err := p.br.GetListOfBooks(context.Background(), domain.GetListOfBooksReq{
		Subject: req.Subject,
	})
	if err != nil {
		return result, err
	}

	for _, item := range res.Books {
		authors := []Author{}
		for _, author := range item.Authors {
			authors = append(authors, Author{
				Name: author.Name,
			})
		}

		result.Books = append(result.Books, Book{
			Key:               item.Key,
			Title:             item.Title,
			EditionCount:      item.EditionCount,
			Authors:           authors,
			LendingIdentifier: item.LendingIdentifier,
		})

	}

	return result, nil
}

func (p bookService) BorrowBook(ctx context.Context, req BorrowBookReq) (BorrowBookRes, error) {
	var result BorrowBookRes

	book, err := p.br.GetBookByKey(ctx, domain.GeBookByKeyReq{
		Key:     req.BookKey,
		Subject: req.Subject,
	})
	if err != nil {
		return result, err
	}

	if err := p.br.BorrowBook(context.Background(), domain.BorrowBookReq{
		Book:       book,
		PickUpDate: req.PickUpDate,
		UserID:     req.UserID,
	}); err != nil {
		return result, err
	}

	authors := []Author{}
	for _, author := range book.Authors {
		authors = append(authors, Author{
			Name: author.Name,
		})
	}

	result = BorrowBookRes{
		Book: Book{
			Key:               book.Key,
			Title:             book.Title,
			EditionCount:      book.EditionCount,
			Authors:           authors,
			LendingIdentifier: book.LendingIdentifier,
		},
		PickUpDate: req.PickUpDate,
		UserID:     req.UserID,
	}

	return result, nil
}

func (p bookService) GetBookReservation(ctx context.Context, req GetBookReservationReq) (map[int][]GetBookReservationRes, error) {
	var result map[int][]GetBookReservationRes = make(map[int][]GetBookReservationRes)

	res, err := p.br.GetBookReservation(ctx, domain.GetBookReservationReq{
		UserID: req.UserID,
	})
	if err != nil {
		return result, err
	}

	for key, value := range res {
		var tmp []GetBookReservationRes
		for _, item := range value {
			tmp = append(tmp, GetBookReservationRes{
				BookKey:    item.Book.Key,
				PickUpDate: item.PickUpDate,
				UserID:     item.UserID,
			})
		}
		result[key] = tmp
	}

	return result, nil
}
