package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"go.bourbon.stream/go-vanityurls-apex/handler"
	yaml "gopkg.in/yaml.v2"
)

func main() {
	addr := ":" + os.Getenv("PORT")

	h, err := handler.NewHandler(fileFetcher{path: "vanity.yaml"})
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", h)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

type fileFetcher struct {
	path string
}

func (ff fileFetcher) Fetch() (*handler.Config, error) {
	var config *handler.Config

	vanity, err := ioutil.ReadFile(ff.path)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(vanity, &config); err != nil {
		return nil, err
	}

	return config, nil
}
