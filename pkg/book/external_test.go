package book

import (
	"bytes"
	"context"
	"errors"
	"net/http/httptest"
	reflect "reflect"
	"testing"

	"gihub.com/gadhittana01/book-project/config"
	"gihub.com/gadhittana01/book-project/pkg/domain"
	gomock "github.com/golang/mock/gomock"
)

func Test_newExternal(t *testing.T) {
	ctrl := gomock.NewController(t)

	httpClientMock := NewMockHttpResource(ctrl)
	cfg := config.GlobalConfig{}

	type args struct {
		cfg        *config.GlobalConfig
		httpclient HttpResource
	}
	tests := []struct {
		name string
		args args
		want external
	}{
		{
			name: "success",
			args: args{
				cfg:        &cfg,
				httpclient: httpClientMock,
			},
			want: &externalModule{
				cfg:        &cfg,
				httpclient: httpClientMock,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newExternal(tt.args.cfg, tt.args.httpclient); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newExternal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getListOfBooks(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()

	type fields struct {
		cfg        *config.GlobalConfig
		httpclient HttpResource
	}
	type args struct {
		ctx context.Context
		req domain.GetListOfBooksReq
	}

	tests := []struct {
		name    string
		fields  func() fields
		args    args
		want    domain.GetListOfBooksResp
		wantErr bool
	}{
		{
			name: "test success",
			args: args{
				ctx: ctx,
				req: domain.GetListOfBooksReq{
					Subject: "love",
				},
			},
			fields: func() fields {
				w := httptest.NewRecorder()
				w.Code = 200
				w.Body = bytes.NewBufferString(`{
					"key": "/subjects/2",
					"name": "2",
					"subject_type": "subject",
					"work_count": 11,
					"works": [
						{
							"key": "/works/OL1908641W",
							"title": "Know Nothing",
							"edition_count": 6,
							"cover_id": 815673,
							"cover_edition_key": "OL965879M",
							"lending_identifier": "knownothingnovel00sett",
							"authors": [
								{
									"key": "/authors/OL228578A",
									"name": "Mary Lee Settle"
								}
							]
						}
					]
				}`)
				res := w.Result()
				defer res.Body.Close()

				httpClientMock := NewMockHttpResource(ctrl)
				httpClientMock.EXPECT().Do(gomock.Any()).Return(res, nil)
				return fields{
					cfg: &config.GlobalConfig{
						BookService: config.BookService{
							Address: "https://dummyaccountsservice.com",
						},
					},
					httpclient: httpClientMock,
				}
			},
			want: domain.GetListOfBooksResp{
				Books: []domain.Book{
					domain.Book{
						Key:          "/works/OL1908641W",
						Title:        "Know Nothing",
						EditionCount: 6,
						Authors: []domain.Author{
							domain.Author{
								Name: "Mary Lee Settle",
							},
						},
						LendingIdentifier: "knownothingnovel00sett",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "test subject empty",
			args: args{
				ctx: ctx,
				req: domain.GetListOfBooksReq{
					Subject: "",
				},
			},
			fields: func() fields {
				httpClientMock := NewMockHttpResource(ctrl)
				return fields{
					cfg: &config.GlobalConfig{
						BookService: config.BookService{
							Address: "https://dummyaccountsservice.com",
						},
					},
					httpclient: httpClientMock,
				}
			},
			want:    domain.GetListOfBooksResp{},
			wantErr: true,
		},
		{
			name: "failed http do",
			args: args{
				ctx: ctx,
				req: domain.GetListOfBooksReq{
					Subject: "love",
				},
			},
			fields: func() fields {
				w := httptest.NewRecorder()
				w.Code = 500
				w.Body = bytes.NewBufferString(`{}`)
				res := w.Result()
				defer res.Body.Close()

				httpClientMock := NewMockHttpResource(ctrl)
				httpClientMock.EXPECT().Do(gomock.Any()).Return(res, errors.New("error"))
				return fields{
					cfg: &config.GlobalConfig{
						BookService: config.BookService{
							Address: "https://dummyaccountsservice.com",
						},
					},
					httpclient: httpClientMock,
				}
			},
			want:    domain.GetListOfBooksResp{},
			wantErr: true,
		},
		{
			name: "error status code",
			args: args{
				ctx: ctx,
				req: domain.GetListOfBooksReq{
					Subject: "love",
				},
			},
			fields: func() fields {
				w := httptest.NewRecorder()
				w.Code = 500
				w.Body = bytes.NewBufferString(`{}`)
				res := w.Result()
				defer res.Body.Close()

				httpClientMock := NewMockHttpResource(ctrl)
				httpClientMock.EXPECT().Do(gomock.Any()).Return(res, nil)
				return fields{
					cfg: &config.GlobalConfig{
						BookService: config.BookService{
							Address: "https://dummyaccountsservice.com",
						},
					},
					httpclient: httpClientMock,
				}
			},
			want:    domain.GetListOfBooksResp{},
			wantErr: true,
		},
		{
			name: "error read resp body",
			args: args{
				ctx: ctx,
				req: domain.GetListOfBooksReq{
					Subject: "love",
				},
			},
			fields: func() fields {
				w := httptest.NewRecorder()
				w.Code = 200
				w.Body = bytes.NewBufferString(``)
				res := w.Result()
				defer res.Body.Close()

				httpClientMock := NewMockHttpResource(ctrl)
				httpClientMock.EXPECT().Do(gomock.Any()).Return(res, nil)
				return fields{
					cfg: &config.GlobalConfig{
						BookService: config.BookService{
							Address: "https://dummyaccountsservice.com",
						},
					},
					httpclient: httpClientMock,
				}
			},
			want:    domain.GetListOfBooksResp{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := tt.fields()
			m := &externalModule{
				cfg:        field.cfg,
				httpclient: field.httpclient,
			}
			got, err := m.getListOfBooks(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("getListOfBooks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getListOfBooks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getBookByKey(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()

	type fields struct {
		cfg        *config.GlobalConfig
		httpclient HttpResource
	}
	type args struct {
		ctx context.Context
		req domain.GeBookByKeyReq
	}

	tests := []struct {
		name    string
		fields  func() fields
		args    args
		want    domain.Book
		wantErr bool
	}{
		{
			name: "test success",
			args: args{
				ctx: ctx,
				req: domain.GeBookByKeyReq{
					Subject: "love",
					Key:     "/works/OL1908641W",
				},
			},
			fields: func() fields {
				w := httptest.NewRecorder()
				w.Code = 200
				w.Body = bytes.NewBufferString(`{
					"key": "/subjects/2",
					"name": "2",
					"subject_type": "subject",
					"work_count": 11,
					"works": [
						{
							"key": "/works/OL1908641W",
							"title": "Know Nothing",
							"edition_count": 6,
							"cover_id": 815673,
							"cover_edition_key": "OL965879M",
							"lending_identifier": "knownothingnovel00sett",
							"authors": [
								{
									"key": "/authors/OL228578A",
									"name": "Mary Lee Settle"
								}
							]
						}
					]
				}`)
				res := w.Result()
				defer res.Body.Close()

				httpClientMock := NewMockHttpResource(ctrl)
				httpClientMock.EXPECT().Do(gomock.Any()).Return(res, nil)
				return fields{
					cfg: &config.GlobalConfig{
						BookService: config.BookService{
							Address: "https://dummyaccountsservice.com",
						},
					},
					httpclient: httpClientMock,
				}
			},
			want: domain.Book{
				Key:          "/works/OL1908641W",
				Title:        "Know Nothing",
				EditionCount: 6,
				Authors: []domain.Author{
					domain.Author{
						Name: "Mary Lee Settle",
					},
				},
				LendingIdentifier: "knownothingnovel00sett",
			},
			wantErr: false,
		},
		{
			name: "test error get list of books",
			args: args{
				ctx: ctx,
				req: domain.GeBookByKeyReq{
					Subject: "love",
					Key:     "/works/OL1908641W",
				},
			},
			fields: func() fields {
				w := httptest.NewRecorder()
				w.Code = 500
				w.Body = bytes.NewBufferString(`{}`)
				res := w.Result()
				defer res.Body.Close()

				httpClientMock := NewMockHttpResource(ctrl)
				httpClientMock.EXPECT().Do(gomock.Any()).Return(res, errors.New("error"))
				return fields{
					cfg: &config.GlobalConfig{
						BookService: config.BookService{
							Address: "https://dummyaccountsservice.com",
						},
					},
					httpclient: httpClientMock,
				}
			},
			want:    domain.Book{},
			wantErr: true,
		},
		{
			name: "test book not found",
			args: args{
				ctx: ctx,
				req: domain.GeBookByKeyReq{
					Subject: "love",
					Key:     "not found",
				},
			},
			fields: func() fields {
				w := httptest.NewRecorder()
				w.Code = 200
				w.Body = bytes.NewBufferString(`{
					"key": "/subjects/2",
					"name": "2",
					"subject_type": "subject",
					"work_count": 11,
					"works": [
						{
							"key": "/works/OL1908641W",
							"title": "Know Nothing",
							"edition_count": 6,
							"cover_id": 815673,
							"cover_edition_key": "OL965879M",
							"lending_identifier": "knownothingnovel00sett",
							"authors": [
								{
									"key": "/authors/OL228578A",
									"name": "Mary Lee Settle"
								}
							]
						}
					]
				}`)
				res := w.Result()
				defer res.Body.Close()

				httpClientMock := NewMockHttpResource(ctrl)
				httpClientMock.EXPECT().Do(gomock.Any()).Return(res, nil)
				return fields{
					cfg: &config.GlobalConfig{
						BookService: config.BookService{
							Address: "https://dummyaccountsservice.com",
						},
					},
					httpclient: httpClientMock,
				}
			},
			want:    domain.Book{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := tt.fields()
			m := &externalModule{
				cfg:        field.cfg,
				httpclient: field.httpclient,
			}
			got, err := m.getBookByKey(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("getBookByKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getBookByKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
