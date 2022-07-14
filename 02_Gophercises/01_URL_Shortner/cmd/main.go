package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/Shresht7/urlshort"
)

func main() {

	filePath := flag.String("path", "paths.yml", "Path to the yml file containing the url paths")

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	bytes, err := os.ReadFile(*filePath)
	if err != nil {
		panic(err)
	}

	var handler http.Handler

	switch path.Ext(*filePath) {
	case "yml", "yaml":
		handler, err = urlshort.YAMLHandler(bytes, mapHandler)
	case "json":
		handler, err = urlshort.JSONHandler(bytes, mapHandler)
	}

	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", handler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
