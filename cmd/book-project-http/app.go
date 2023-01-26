package main

import (
	"gihub.com/gadhittana01/book-project/config"
	"gihub.com/gadhittana01/book-project/handler/resthttp"
	"gihub.com/gadhittana01/book-project/pkg/book"
	httpClient "gihub.com/gadhittana01/book-project/pkg/http_client"
	"gihub.com/gadhittana01/book-project/services"
)

func initApp(c *config.GlobalConfig) error {
	httpClient := httpClient.New(httpClient.HttpClientDep{
		Config: c,
	})
	bookPkg := book.New(c, httpClient)

	bs, err := services.NewBookService(services.BookDependencies{
		BR: bookPkg,
	})
	if err != nil {
		return err
	}

	return startHTTPServer(resthttp.NewRoutes(resthttp.RouterDependencies{
		BS: bs,
	}), c)
}
