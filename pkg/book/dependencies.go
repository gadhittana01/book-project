package book

import (
	"net/http"
)

type (
	HttpResource interface {
		Do(req *http.Request) (*http.Response, error)
	}
)
