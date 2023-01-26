package resthttp

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"gihub.com/gadhittana01/book-project/services"
	"github.com/golang/mock/gomock"
)

func Test_NewBookService(t *testing.T) {
	ctrl := gomock.NewController(t)

	bookMock := NewMockBookService(ctrl)
	type args struct {
		service BookService
	}
	tests := []struct {
		name string
		args args
		want *bookHandler
	}{
		{
			args: args{
				service: bookMock,
			},
			want: &bookHandler{
				service: bookMock,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newBookHandler(tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newBookHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GetListOfBooks(t *testing.T) {
	ctrl := gomock.NewController(t)
	sampleReq := httptest.NewRequest("GET", "http://localhost:8000/get-books?subject=love", strings.NewReader(""))
	sampleResp := httptest.NewRecorder()

	badReq := httptest.NewRequest("GET", "http://localhost:8000/get-books", strings.NewReader(""))
	badResp := httptest.NewRecorder()

	type fields struct {
		service BookService
	}
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name   string
		fields func() fields
		args   args
	}{
		{
			name: "test normal flow",
			fields: func() fields {
				bookMock := NewMockBookService(ctrl)
				bookMock.EXPECT().GetListOfBooks(gomock.Any(), services.GetListOfBooksReq{
					Subject: "love",
				}).Return(services.GetListOfBooksResp{
					Books: []services.Book{
						services.Book{
							Key:          "123",
							Title:        "ABC",
							EditionCount: 123,
							Authors: []services.Author{
								services.Author{
									Name: "Jack",
								},
							},
							LendingIdentifier: "TST",
						},
					},
				}, nil)
				return fields{
					service: bookMock,
				}
			},
			args: args{
				w:   sampleResp,
				req: sampleReq,
			},
		},
		{
			name: "test bad request",
			fields: func() fields {
				bookMock := NewMockBookService(ctrl)
				return fields{
					service: bookMock,
				}
			},
			args: args{
				w:   badResp,
				req: badReq,
			},
		},
		{
			name: "test internal server error",
			fields: func() fields {
				bookMock := NewMockBookService(ctrl)
				bookMock.EXPECT().GetListOfBooks(gomock.Any(), services.GetListOfBooksReq{
					Subject: "love",
				}).Return(services.GetListOfBooksResp{}, errors.New("error"))
				return fields{
					service: bookMock,
				}
			},
			args: args{
				w:   sampleResp,
				req: sampleReq,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := tt.fields()
			i := bookHandler{
				service: field.service,
			}
			i.GetListOfBooks(tt.args.w, tt.args.req)
		})
	}
}

func Test_BorrowBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	sampleReq := httptest.NewRequest("GET", "http://localhost:8000/borrow-book", strings.NewReader(`{
		"key" : "/works/OL98501W",
		"pickup_date" : "2022-02-26",
		"subject" : "love",
		"user_id" : 2
	}`))
	sampleResp := httptest.NewRecorder()

	internalErrReq := httptest.NewRequest("GET", "http://localhost:8000/borrow-book", strings.NewReader(`{
		"key" : "/works/OL98501W",
		"pickup_date" : "2022-02-26",
		"subject" : "love",
		"user_id" : 2
	}`))
	internalErrResp := httptest.NewRecorder()

	badReq := httptest.NewRequest("GET", "http://localhost:8000/borrow-book", strings.NewReader(""))
	badResp := httptest.NewRecorder()

	type fields struct {
		service BookService
	}
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name   string
		fields func() fields
		args   args
	}{
		{
			name: "test normal flow",
			fields: func() fields {
				bookMock := NewMockBookService(ctrl)
				bookMock.EXPECT().BorrowBook(gomock.Any(), services.BorrowBookReq{
					BookKey:    "/works/OL98501W",
					PickUpDate: "2022-02-26",
					Subject:    "love",
					UserID:     2,
				}).Return(services.BorrowBookRes{
					Book: services.Book{
						Key:          "/works/OL98501W",
						Title:        "Hello World",
						EditionCount: 12,
						Authors: []services.Author{
							services.Author{
								Name: "test",
							},
						},
						LendingIdentifier: "1234",
					},
				}, nil)
				return fields{
					service: bookMock,
				}
			},
			args: args{
				w:   sampleResp,
				req: sampleReq,
			},
		},
		{
			name: "test bad request",
			fields: func() fields {
				bookMock := NewMockBookService(ctrl)
				return fields{
					service: bookMock,
				}
			},
			args: args{
				w:   badResp,
				req: badReq,
			},
		},
		{
			name: "test internal server error",
			fields: func() fields {
				bookMock := NewMockBookService(ctrl)
				bookMock.EXPECT().BorrowBook(gomock.Any(), services.BorrowBookReq{
					BookKey:    "/works/OL98501W",
					PickUpDate: "2022-02-26",
					Subject:    "love",
					UserID:     2,
				}).Return(services.BorrowBookRes{}, errors.New("error"))
				return fields{
					service: bookMock,
				}
			},
			args: args{
				w:   internalErrResp,
				req: internalErrReq,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := tt.fields()
			i := bookHandler{
				service: field.service,
			}
			i.BorrowBook(tt.args.w, tt.args.req)
		})
	}
}

func Test_GetBookReservation(t *testing.T) {
	ctrl := gomock.NewController(t)
	sampleReq := httptest.NewRequest("GET", "http://localhost:8000/get-book-reservation", strings.NewReader(""))
	sampleResp := httptest.NewRecorder()

	sampleUidReq := httptest.NewRequest("GET", "http://localhost:8000/get-book-reservation?user_id=1", strings.NewReader(""))
	sampleUidResp := httptest.NewRecorder()

	internalErrReq := httptest.NewRequest("GET", "http://localhost:8000/get-book-reservation?user_id=1", strings.NewReader(""))
	internalErrResp := httptest.NewRecorder()

	badReq := httptest.NewRequest("GET", "http://localhost:8000/get-book-reservation?user_id=boba", strings.NewReader(""))
	badResp := httptest.NewRecorder()

	type fields struct {
		service BookService
	}
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name   string
		fields func() fields
		args   args
	}{
		{
			name: "test normal flow without user id",
			fields: func() fields {
				bookMock := NewMockBookService(ctrl)
				bookMock.EXPECT().GetBookReservation(gomock.Any(), services.GetBookReservationReq{
					UserID: 0,
				}).Return(map[int][]services.GetBookReservationRes{
					1: []services.GetBookReservationRes{
						services.GetBookReservationRes{
							BookKey:    "123",
							PickUpDate: "2022-10-10",
							UserID:     1,
						},
					},
				}, nil)
				return fields{
					service: bookMock,
				}
			},
			args: args{
				w:   sampleResp,
				req: sampleReq,
			},
		},
		{
			name: "test normal flow with user id",
			fields: func() fields {
				bookMock := NewMockBookService(ctrl)
				bookMock.EXPECT().GetBookReservation(gomock.Any(), services.GetBookReservationReq{
					UserID: 1,
				}).Return(map[int][]services.GetBookReservationRes{
					1: []services.GetBookReservationRes{
						services.GetBookReservationRes{
							BookKey:    "123",
							PickUpDate: "2022-10-10",
							UserID:     1,
						},
					},
				}, nil)
				return fields{
					service: bookMock,
				}
			},
			args: args{
				w:   sampleUidResp,
				req: sampleUidReq,
			},
		},
		{
			name: "test bad request",
			fields: func() fields {
				bookMock := NewMockBookService(ctrl)
				return fields{
					service: bookMock,
				}
			},
			args: args{
				w:   badResp,
				req: badReq,
			},
		},
		{
			name: "test internal server error",
			fields: func() fields {
				bookMock := NewMockBookService(ctrl)
				bookMock.EXPECT().GetBookReservation(gomock.Any(), services.GetBookReservationReq{
					UserID: 1,
				}).Return(map[int][]services.GetBookReservationRes{}, errors.New("error"))
				return fields{
					service: bookMock,
				}
			},
			args: args{
				w:   internalErrResp,
				req: internalErrReq,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := tt.fields()
			i := bookHandler{
				service: field.service,
			}
			i.GetBookReservation(tt.args.w, tt.args.req)
		})
	}
}
