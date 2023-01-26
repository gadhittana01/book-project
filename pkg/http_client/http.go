package httpclient

import (
	"net/http"
	"time"

	"gihub.com/gadhittana01/book-project/config"
)

type HttpClientDep struct {
	Config *config.GlobalConfig
}

type httpModule struct {
	client *http.Client
}

type HttpIFace interface {
	Do(req *http.Request) (*http.Response, error)
}

func New(dep HttpClientDep) HttpIFace {
	client := http.Client{
		Timeout: time.Duration(dep.Config.HttpClientConfig.TimeoutMS) * time.Millisecond,
		Transport: &http.Transport{
			MaxIdleConns:        dep.Config.HttpClientConfig.MaxIdleConns,
			MaxIdleConnsPerHost: dep.Config.HttpClientConfig.MaxIdleConnsPerHost,
			MaxConnsPerHost:     dep.Config.HttpClientConfig.MaxConnsPerHost,
			IdleConnTimeout:     time.Duration(dep.Config.HttpClientConfig.IdleConnTimeoutSec) * time.Second,
		},
	}

	return &httpModule{
		client: &client,
	}
}

func (hm *httpModule) Do(req *http.Request) (*http.Response, error) {
	return hm.client.Do(req)
}
