package main

import (
	"fireBackend/cmd/webapp/routes"

	"github.com/MakMoinee/go-mith/pkg/goserve"
)

func main() {
	newHttpService := goserve.NewService(":8443")
	routes.Set(newHttpService)

	if err := newHttpService.Start(); err != nil {
		panic(err)
	}
}
