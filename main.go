package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"go.bourbon.stream/go-vanityurls-apex/handler"
)

func main() {
	configPath := "vanity.yaml"
	addr := ":" + os.Getenv("PORT")
	vanity, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}
	h, err := handler.NewHandler(vanity)
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", h)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
