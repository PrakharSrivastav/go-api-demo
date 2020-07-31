package main

import (
	"github.com/PrakharSrivastav/go-api-demo/pkg"
	"log"
	"net/http"
)

func main() {
	app := pkg.New()
	r := app.Routes()
	err := http.ListenAndServe(":3000", r)
	log.Fatal(err)
}
