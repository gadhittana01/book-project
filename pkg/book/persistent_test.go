package book

import (
	"context"
	reflect "reflect"
	"testing"

	"gihub.com/gadhittana01/book-project/pkg/domain"
)

func Test_newPersistent(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		args args
		want persistent
	}{
		{
			name: "success",
			args: args{},
			want: &persistentModule{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newPersistent(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newPersistent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_borrowBook(t *testing.T) {
	ctx := context.Background()

	type fields struct {
	}
	type args struct {
		ctx context.Context
		req domain.BorrowBookReq
	}

	tests := []struct {
		name    string
		fields  func() fields
		args    args
		wantErr bool
	}{
		{
			name: "test success",
			args: args{
				ctx: ctx,
				req: domain.BorrowBookReq{
					Book: domain.Book{
						Key:          "/works/OL98501W",
						Title:        "test",
						EditionCount: 123,
						Authors: []domain.Author{
							domain.Author{
								Name: "Giri Putra Adhittana",
							},
						},
						LendingIdentifier: "123",
					},
					UserID:     1,
					PickUpDate: "2022-01-25",
				},
			},
			fields: func() fields {
				return fields{}
			},
			wantErr: false,
		},
		{
			name: "test user id is empty",
			args: args{
				ctx: ctx,
				req: domain.BorrowBookReq{
					Book: domain.Book{
						Key:          "/works/OL98501W",
						Title:        "test",
						EditionCount: 123,
						Authors: []domain.Author{
							domain.Author{
								Name: "Giri Putra Adhittana",
							},
						},
						LendingIdentifier: "123",
					},
					UserID:     0,
					PickUpDate: "2022-01-25",
				},
			},
			fields: func() fields {
				return fields{}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &persistentModule{}
			err := m.borrowBook(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("borrowBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_getBookReservation(t *testing.T) {
	ctx := context.Background()
	const userID = 1

	type fields struct {
	}
	type args struct {
		ctx context.Context
		req domain.GetBookReservationReq
	}

	tests := []struct {
		name    string
		fields  func() fields
		args    args
		want    map[int][]domain.BorrowBookReq
		wantErr bool
	}{
		{
			name: "test success",
			args: args{
				ctx: ctx,
				req: domain.GetBookReservationReq{
					UserID: 1,
				},
			},
			fields: func() fields {
				return fields{}
			},
			want: map[int][]domain.BorrowBookReq{
				userID: books[userID],
			},
			wantErr: false,
		},
		{
			name: "test success empty userID",
			args: args{
				ctx: ctx,
				req: domain.GetBookReservationReq{
					UserID: 0,
				},
			},
			fields: func() fields {
				return fields{}
			},
			want:    books,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &persistentModule{}
			got, err := m.getBookReservation(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("getBookReservation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getBookReservation() = %v, want %v", got, tt.want)
			}
		})
	}
}
