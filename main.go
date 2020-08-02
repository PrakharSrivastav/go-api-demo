package main

import (
	"github.com/PrakharSrivastav/go-api-demo/pkg"
	"log"
	"net/http"
)

func main() {
	app, err := pkg.New()
	if err != nil {
		log.Fatal(err)
	}
	r := app.Routes()
	err = http.ListenAndServe(":3000", r)
	log.Fatal(err)
}
