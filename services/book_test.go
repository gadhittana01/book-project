package services

import (
	"context"
	"errors"
	reflect "reflect"
	"testing"

	"gihub.com/gadhittana01/book-project/pkg/domain"
	gomock "github.com/golang/mock/gomock"
)

func Test_NewBookService(t *testing.T) {
	ctrl := gomock.NewController(t)
	bookMock := NewMockBookResource(ctrl)

	type args struct {
		dep BookDependencies
	}
	tests := []struct {
		name    string
		args    args
		want    BookService
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				dep: BookDependencies{
					BR: bookMock,
				},
			},
			want: &bookService{
				br: bookMock,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBookService(tt.args.dep)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBookService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBookService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GetListOfBooks(t *testing.T) {
	ctrl := gomock.NewController(t)

	type args struct {
		ctx context.Context
		req GetListOfBooksReq
	}
	tests := []struct {
		name    string
		fields  func() bookService
		args    args
		want    GetListOfBooksResp
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: GetListOfBooksReq{
					Subject: "love",
				},
			},
			fields: func() bookService {
				bookMock := NewMockBookResource(ctrl)
				bookMock.EXPECT().GetListOfBooks(gomock.Any(), domain.GetListOfBooksReq{
					Subject: "love",
				}).Return(domain.GetListOfBooksResp{
					Books: []domain.Book{
						domain.Book{
							Key:          "123",
							Title:        "hello",
							EditionCount: 123,
							Authors: []domain.Author{
								domain.Author{
									Name: "boba",
								},
							},
							LendingIdentifier: "456",
						},
					},
				}, nil)

				return bookService{
					br: bookMock,
				}
			},
			want: GetListOfBooksResp{
				Books: []Book{
					Book{
						Key:          "123",
						Title:        "hello",
						EditionCount: 123,
						Authors: []Author{
							Author{
								Name: "boba",
							},
						},
						LendingIdentifier: "456",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "error get list of books",
			args: args{
				ctx: context.Background(),
				req: GetListOfBooksReq{
					Subject: "love",
				},
			},
			fields: func() bookService {
				bookMock := NewMockBookResource(ctrl)
				bookMock.EXPECT().GetListOfBooks(gomock.Any(), domain.GetListOfBooksReq{
					Subject: "love",
				}).Return(domain.GetListOfBooksResp{}, errors.New("error"))

				return bookService{
					br: bookMock,
				}
			},
			want:    GetListOfBooksResp{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.fields()
			got, err := m.GetListOfBooks(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetListOfBooks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetListOfBooks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_BorrowBook(t *testing.T) {
	ctrl := gomock.NewController(t)

	type args struct {
		ctx context.Context
		req BorrowBookReq
	}
	tests := []struct {
		name    string
		fields  func() bookService
		args    args
		want    BorrowBookRes
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: BorrowBookReq{
					BookKey:    "123",
					PickUpDate: "2022-01-01",
					Subject:    "love",
					UserID:     1,
				},
			},
			fields: func() bookService {
				bookMock := NewMockBookResource(ctrl)
				bookMock.EXPECT().GetBookByKey(gomock.Any(), domain.GeBookByKeyReq{
					Key:     "123",
					Subject: "love",
				}).Return(domain.Book{
					Key:          "123",
					Title:        "hello",
					EditionCount: 123,
					Authors: []domain.Author{
						domain.Author{
							Name: "boba",
						},
					},
					LendingIdentifier: "456",
				}, nil)

				bookMock.EXPECT().BorrowBook(gomock.Any(), domain.BorrowBookReq{
					Book: domain.Book{
						Key:          "123",
						Title:        "hello",
						EditionCount: 123,
						Authors: []domain.Author{
							domain.Author{
								Name: "boba",
							},
						},
						LendingIdentifier: "456",
					},
					PickUpDate: "2022-01-01",
					UserID:     1,
				}).Return(nil)

				return bookService{
					br: bookMock,
				}
			},
			want: BorrowBookRes{
				Book: Book{
					Key:          "123",
					Title:        "hello",
					EditionCount: 123,
					Authors: []Author{
						Author{
							Name: "boba",
						},
					},
					LendingIdentifier: "456",
				},
				PickUpDate: "2022-01-01",
				UserID:     1,
			},
			wantErr: false,
		},
		{
			name: "book not found",
			args: args{
				ctx: context.Background(),
				req: BorrowBookReq{
					BookKey:    "123",
					PickUpDate: "2022-01-01",
					Subject:    "love",
					UserID:     1,
				},
			},
			fields: func() bookService {
				bookMock := NewMockBookResource(ctrl)
				bookMock.EXPECT().GetBookByKey(gomock.Any(), domain.GeBookByKeyReq{
					Key:     "123",
					Subject: "love",
				}).Return(domain.Book{}, errors.New("book not found"))

				return bookService{
					br: bookMock,
				}
			},
			want:    BorrowBookRes{},
			wantErr: true,
		},
		{
			name: "error borrow book",
			args: args{
				ctx: context.Background(),
				req: BorrowBookReq{
					BookKey:    "123",
					PickUpDate: "2022-01-01",
					Subject:    "love",
					UserID:     1,
				},
			},
			fields: func() bookService {
				bookMock := NewMockBookResource(ctrl)
				bookMock.EXPECT().GetBookByKey(gomock.Any(), domain.GeBookByKeyReq{
					Key:     "123",
					Subject: "love",
				}).Return(domain.Book{
					Key:          "123",
					Title:        "hello",
					EditionCount: 123,
					Authors: []domain.Author{
						domain.Author{
							Name: "boba",
						},
					},
					LendingIdentifier: "456",
				}, nil)

				bookMock.EXPECT().BorrowBook(gomock.Any(), domain.BorrowBookReq{
					Book: domain.Book{
						Key:          "123",
						Title:        "hello",
						EditionCount: 123,
						Authors: []domain.Author{
							domain.Author{
								Name: "boba",
							},
						},
						LendingIdentifier: "456",
					},
					PickUpDate: "2022-01-01",
					UserID:     1,
				}).Return(errors.New("error"))

				return bookService{
					br: bookMock,
				}
			},
			want:    BorrowBookRes{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.fields()
			got, err := m.BorrowBook(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("BorrowBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BorrowBook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GetBookReservation(t *testing.T) {
	ctrl := gomock.NewController(t)

	type args struct {
		ctx context.Context
		req GetBookReservationReq
	}
	tests := []struct {
		name    string
		fields  func() bookService
		args    args
		want    map[int][]GetBookReservationRes
		wantErr bool
	}{
		{
			name: "success with user id",
			args: args{
				ctx: context.Background(),
				req: GetBookReservationReq{
					UserID: 1,
				},
			},
			fields: func() bookService {
				bookMock := NewMockBookResource(ctrl)
				bookMock.EXPECT().GetBookReservation(gomock.Any(), domain.GetBookReservationReq{
					UserID: 1,
				}).Return(map[int][]domain.BorrowBookReq{
					1: []domain.BorrowBookReq{
						domain.BorrowBookReq{
							Book: domain.Book{
								Key:          "123",
								Title:        "ABC",
								EditionCount: 123,
								Authors: []domain.Author{
									domain.Author{
										Name: "Jack",
									},
								},
								LendingIdentifier: "TST",
							},
							PickUpDate: "2022-01-01",
							UserID:     1,
						},
					},
				}, nil)

				return bookService{
					br: bookMock,
				}
			},
			want: map[int][]GetBookReservationRes{
				1: []GetBookReservationRes{
					GetBookReservationRes{
						BookKey:    "123",
						PickUpDate: "2022-01-01",
						UserID:     1,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "success without user id",
			args: args{
				ctx: context.Background(),
				req: GetBookReservationReq{
					UserID: 0,
				},
			},
			fields: func() bookService {
				bookMock := NewMockBookResource(ctrl)
				bookMock.EXPECT().GetBookReservation(gomock.Any(), domain.GetBookReservationReq{
					UserID: 0,
				}).Return(map[int][]domain.BorrowBookReq{
					1: []domain.BorrowBookReq{
						domain.BorrowBookReq{
							Book: domain.Book{
								Key:          "123",
								Title:        "ABC",
								EditionCount: 123,
								Authors: []domain.Author{
									domain.Author{
										Name: "Jack",
									},
								},
								LendingIdentifier: "TST",
							},
							PickUpDate: "2022-01-01",
							UserID:     1,
						},
					},
					2: []domain.BorrowBookReq{
						domain.BorrowBookReq{
							Book: domain.Book{
								Key:          "123",
								Title:        "ABC",
								EditionCount: 123,
								Authors: []domain.Author{
									domain.Author{
										Name: "Jack",
									},
								},
								LendingIdentifier: "TST",
							},
							PickUpDate: "2022-01-01",
							UserID:     1,
						},
					},
				}, nil)

				return bookService{
					br: bookMock,
				}
			},
			want: map[int][]GetBookReservationRes{
				1: []GetBookReservationRes{
					GetBookReservationRes{
						BookKey:    "123",
						PickUpDate: "2022-01-01",
						UserID:     1,
					},
				},
				2: []GetBookReservationRes{
					GetBookReservationRes{
						BookKey:    "123",
						PickUpDate: "2022-01-01",
						UserID:     1,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "error get book reservation",
			args: args{
				ctx: context.Background(),
				req: GetBookReservationReq{
					UserID: 1,
				},
			},
			fields: func() bookService {
				bookMock := NewMockBookResource(ctrl)
				bookMock.EXPECT().GetBookReservation(gomock.Any(), domain.GetBookReservationReq{
					UserID: 1,
				}).Return(map[int][]domain.BorrowBookReq{}, errors.New("error"))

				return bookService{
					br: bookMock,
				}
			},
			want:    map[int][]GetBookReservationRes{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.fields()
			got, err := m.GetBookReservation(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBookReservation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBookReservation() = %v, want %v", got, tt.want)
			}
		})
	}
}
