package book

import (
	"context"
	reflect "reflect"
	"testing"

	"gihub.com/gadhittana01/book-project/config"
	"gihub.com/gadhittana01/book-project/pkg/domain"
	gomock "github.com/golang/mock/gomock"
)

func TestNew(t *testing.T) {
	ctrl := gomock.NewController(t)

	httpMock := NewMockHttpResource(ctrl)
	cfg := config.GlobalConfig{}

	type args struct {
		cfg        *config.GlobalConfig
		httpclient HttpResource
	}
	tests := []struct {
		name    string
		args    args
		want    *module
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				cfg:        &cfg,
				httpclient: httpMock,
			},
			want: &module{
				external: newExternal(
					&cfg,
					httpMock,
				),
				persistent: newPersistent(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.cfg, tt.args.httpclient); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GetListOfBooks(t *testing.T) {
	ctrl := gomock.NewController(t)

	type args struct {
		ctx context.Context
		req domain.GetListOfBooksReq
	}
	tests := []struct {
		name    string
		mock    func() *module
		args    args
		want    domain.GetListOfBooksResp
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: domain.GetListOfBooksReq{
					Subject: "love",
				},
			},
			mock: func() *module {
				extMock := NewMockexternal(ctrl)
				pstMock := NewMockpersistent(ctrl)

				extMock.EXPECT().getListOfBooks(gomock.Any(), domain.GetListOfBooksReq{
					Subject: "love",
				}).Return(domain.GetListOfBooksResp{
					Books: []domain.Book{
						domain.Book{
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
					},
				}, nil)

				return &module{
					external:   extMock,
					persistent: pstMock,
				}
			},
			want: domain.GetListOfBooksResp{
				Books: []domain.Book{
					domain.Book{
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
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.mock()
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
		req domain.BorrowBookReq
	}
	tests := []struct {
		name    string
		mock    func() *module
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: domain.BorrowBookReq{
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
			mock: func() *module {
				extMock := NewMockexternal(ctrl)
				pstMock := NewMockpersistent(ctrl)
				pstMock.EXPECT().borrowBook(gomock.Any(), domain.BorrowBookReq{
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
				}).Return(nil)

				return &module{
					external:   extMock,
					persistent: pstMock,
				}
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.mock()
			err := m.BorrowBook(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("BorrowBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_GetBookByKey(t *testing.T) {
	ctrl := gomock.NewController(t)

	type args struct {
		ctx context.Context
		req domain.GeBookByKeyReq
	}
	tests := []struct {
		name    string
		mock    func() *module
		args    args
		want    domain.Book
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: domain.GeBookByKeyReq{
					Subject: "love",
					Key:     "123",
				},
			},
			mock: func() *module {
				extMock := NewMockexternal(ctrl)
				pstMock := NewMockpersistent(ctrl)

				extMock.EXPECT().getBookByKey(gomock.Any(), domain.GeBookByKeyReq{
					Subject: "love",
					Key:     "123",
				}).Return(domain.Book{
					Key:          "123",
					Title:        "ABC",
					EditionCount: 123,
					Authors: []domain.Author{
						domain.Author{
							Name: "Jack",
						},
					},
					LendingIdentifier: "TST",
				}, nil)

				return &module{
					external:   extMock,
					persistent: pstMock,
				}
			},
			want: domain.Book{
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
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.mock()
			got, err := m.GetBookByKey(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBookByKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBookByKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GetBookReservation(t *testing.T) {
	ctrl := gomock.NewController(t)

	type args struct {
		ctx context.Context
		req domain.GetBookReservationReq
	}
	tests := []struct {
		name    string
		mock    func() *module
		args    args
		want    map[int][]domain.BorrowBookReq
		wantErr bool
	}{
		{
			name: "success get book reservation by ID",
			args: args{
				ctx: context.Background(),
				req: domain.GetBookReservationReq{
					UserID: 1,
				},
			},
			mock: func() *module {
				extMock := NewMockexternal(ctrl)
				pstMock := NewMockpersistent(ctrl)

				pstMock.EXPECT().getBookReservation(gomock.Any(), domain.GetBookReservationReq{
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
						},
					},
				}, nil)

				return &module{
					external:   extMock,
					persistent: pstMock,
				}
			},
			want: map[int][]domain.BorrowBookReq{
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
					},
				},
			},
			wantErr: false,
		},
		{
			name: "success get book reservation without ID",
			args: args{
				ctx: context.Background(),
				req: domain.GetBookReservationReq{
					UserID: 0,
				},
			},
			mock: func() *module {
				extMock := NewMockexternal(ctrl)
				pstMock := NewMockpersistent(ctrl)

				pstMock.EXPECT().getBookReservation(gomock.Any(), domain.GetBookReservationReq{
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
						},
					},
				}, nil)

				return &module{
					external:   extMock,
					persistent: pstMock,
				}
			},
			want: map[int][]domain.BorrowBookReq{
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
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.mock()
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
