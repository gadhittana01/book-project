package main

import (
	"log"

	"gihub.com/gadhittana01/book-project/config"
	"gihub.com/gadhittana01/book-project/helper"
)

func main() {
	config := &config.GlobalConfig{}
	helper.LoadConfig(config)
	err := initApp(config)
	if err != nil {
		log.Println(err)
	}
}
