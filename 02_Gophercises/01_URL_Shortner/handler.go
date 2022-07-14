package urlshort

import (
	"encoding/json"
	"net/http"

	"github.com/go-yaml/yaml"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url, ok := pathsToUrls[r.RequestURI]
		if ok {
			http.Redirect(w, r, url, http.StatusTemporaryRedirect)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

type URLLinks struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func buildMap(y []URLLinks) map[string]string {
	var m = map[string]string{}
	for _, entry := range y {
		m[entry.Path] = entry.Url
	}
	return m
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var parsedYml []URLLinks
	err := yaml.Unmarshal(yml, &parsedYml)
	if err != nil {
		return nil, err
	}
	urlMap := buildMap(parsedYml)
	handler := MapHandler(urlMap, fallback)
	return handler, nil
}

func JSONHandler(jsn []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var parsedJSON []URLLinks
	err := json.Unmarshal(jsn, &parsedJSON)
	if err != nil {
		return nil, err
	}
	urlMap := buildMap(parsedJSON)
	handler := MapHandler(urlMap, fallback)
	return handler, nil
}
