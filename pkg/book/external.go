package book

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"gihub.com/gadhittana01/book-project/config"
	"gihub.com/gadhittana01/book-project/pkg/domain"
)

type external interface {
	getListOfBooks(ctx context.Context, req domain.GetListOfBooksReq) (domain.GetListOfBooksResp, error)
	getBookByKey(ctx context.Context, req domain.GeBookByKeyReq) (domain.Book, error)
}

type externalModule struct {
	cfg        *config.GlobalConfig
	httpclient HttpResource
}

func newExternal(cfg *config.GlobalConfig, httpclient HttpResource) external {
	return &externalModule{
		cfg:        cfg,
		httpclient: httpclient,
	}
}

const (
	PathGetListOfBooks = "/subjects"
)

func (m *externalModule) getListOfBooks(ctx context.Context, req domain.GetListOfBooksReq) (domain.GetListOfBooksResp, error) {
	res := domain.GetListOfBooksResp{}

	if req.Subject == "" {
		return res, errors.New("Subject cannot be empty")
	}

	URL := m.cfg.BookService.Address + PathGetListOfBooks + "/" + req.Subject + ".json"

	reqHttp, err := http.NewRequestWithContext(ctx, http.MethodGet, URL, nil)
	if err != nil {
		return res, err
	}

	resHttp, err := m.httpclient.Do(reqHttp)
	if err != nil {
		return res, err
	}

	defer resHttp.Body.Close()

	if resHttp.StatusCode != 200 {
		return res, errors.New("error external call")
	}

	resBody, err := ioutil.ReadAll(resHttp.Body)
	if err != nil {
		return res, err
	}

	if err = json.Unmarshal(resBody, &res); err != nil {
		return res, err
	}

	return res, nil
}

func (m *externalModule) getBookByKey(ctx context.Context, req domain.GeBookByKeyReq) (domain.Book, error) {
	res := domain.Book{}
	resBooks, err := m.getListOfBooks(ctx, domain.GetListOfBooksReq{
		Subject: req.Subject,
	})
	if err != nil {
		return res, err
	}

	for _, item := range resBooks.Books {
		if item.Key == req.Key {
			return item, nil
		}
	}

	return res, errors.New("Book not found")
}
