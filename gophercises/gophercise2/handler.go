package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

type pathToURL struct {
	Path string
	URL  string
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		shortURL := r.URL.Path
		if longURL, ok := pathsToUrls[shortURL]; ok {
			http.Redirect(rw, r, longURL, http.StatusPermanentRedirect)
		}
		fallback.ServeHTTP(rw, r)
	})
}

func parseYAML(yamlData []byte) ([]pathToURL, error) {
	pathsToUrls := make([]pathToURL, 0)
	err := yaml.Unmarshal(yamlData, &pathsToUrls)
	return pathsToUrls, err
}

func parseJSON(jsonData []byte) ([]pathToURL, error) {
	pathsToUrls := make([]pathToURL, 0)
	err := json.Unmarshal(jsonData, &pathsToUrls)
	return pathsToUrls, err
}

func buildMap(pathToUrls []pathToURL) map[string]string {
	urlsMap := make(map[string]string)
	for _, pathToUrl := range pathToUrls {
		urlsMap[pathToUrl.Path] = pathToUrl.URL
	}
	return urlsMap
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
func YAMLHandler(yamlData []byte, fallback http.Handler) http.HandlerFunc {
	pathToUrls, err := parseYAML(yamlData)
	if err != nil {
		exit(fmt.Sprintf("Unable to parse the YAML data. Error: %v", err))
	}
	urlsMap := buildMap(pathToUrls)
	return MapHandler(urlsMap, fallback)
}

func JSONHandler(jsonData []byte, fallback http.Handler) http.HandlerFunc {
	pathToUrls, err := parseJSON(jsonData)
	if err != nil {
		exit(fmt.Sprintf("Unable to parse the JSON data. Error: %v", err))
	}
	urlsMap := buildMap(pathToUrls)
	return MapHandler(urlsMap, fallback)
}
