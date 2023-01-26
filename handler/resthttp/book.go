package resthttp

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"gihub.com/gadhittana01/book-project/services"
)

type bookHandler struct {
	service BookService
}

func newBookHandler(service BookService) *bookHandler {
	return &bookHandler{
		service: service,
	}
}
func (p bookHandler) GetListOfBooks(w http.ResponseWriter, r *http.Request) {
	resp := newResponse(time.Now())

	subject := r.URL.Query().Get("subject")
	if subject == "" {
		resp.setBadRequest("Invalid Request Parameter", w)
		return
	}
	res, err := p.service.GetListOfBooks(context.Background(), services.GetListOfBooksReq{
		Subject: subject,
	})
	if err != nil {
		resp.setInternalServerError(err.Error(), w)
		return
	}

	resp.setOK(map[string]interface{}{
		"books": res.Books,
	}, w)
	return
}

func (p bookHandler) BorrowBook(w http.ResponseWriter, r *http.Request) {
	resp := newResponse(time.Now())

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resp.setBadRequest(err.Error(), w)
		return
	}

	reqBody := services.BorrowBookReq{}
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		resp.setBadRequest(err.Error(), w)
		return
	}
	res, err := p.service.BorrowBook(context.Background(), reqBody)
	if err != nil {
		resp.setInternalServerError(err.Error(), w)
		return
	}

	resp.setOK(map[string]interface{}{
		"data": fmt.Sprintf("Book with key %s successfully reserved at %s", res.Book.Key, res.PickUpDate),
	}, w)
	return
}

func (p bookHandler) GetBookReservation(w http.ResponseWriter, r *http.Request) {
	var (
		uid int
		err error
	)

	resp := newResponse(time.Now())

	query := r.URL.Query()
	userIDString := strings.TrimSpace(query.Get("user_id"))
	if userIDString != "" {
		uid, err = strconv.Atoi(userIDString)
		if err != nil {
			resp.setBadRequest(err.Error(), w)
			return
		}
	}

	res, err := p.service.GetBookReservation(context.Background(), services.GetBookReservationReq{
		UserID: uid,
	})
	if err != nil {
		resp.setInternalServerError(err.Error(), w)
		return
	}

	resp.setOK(map[string]interface{}{
		"data": res,
	}, w)
	return
}
