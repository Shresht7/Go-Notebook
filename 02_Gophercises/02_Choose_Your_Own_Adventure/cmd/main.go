package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	cyoa "github.com/Shresht7/Gophercises/ChooseYourOwnAdventure"
)

func main() {
	port := flag.Int("port", 8080, "the port to run on")
	fileName := flag.String("file", "gopher.json", "path to the story")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.StoryFromJSON(file)
	if err != nil {
		panic(err)
	}

	handler := cyoa.NewHandler(story, nil)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), handler); err != nil {
		panic(err)
	}
}
