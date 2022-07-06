package main

import (
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/dogs" {
		http.Redirect(w, r, "https://google.com", http.StatusFound)
	}
}

func parseYaml(yamlFile []byte) (map[string]string, error) {
	result := make(map[string]string)
	yaml.Unmarshal(yamlFile, &result)
	return result, nil
}

func main() {
	file, _ := os.ReadFile("map.yaml")
	handler, _ := YAMLHandler(file)
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func MapHandler(data map[string]string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if data[r.URL.Path] != "" {
			http.Redirect(w, r, data[r.URL.Path], 301)
		}
	})
}

func YAMLHandler(yaml []byte) (http.HandlerFunc, error) {
	parsedYaml, err := parseYaml(yaml)
	if err != nil {
		return nil, err
	}
	return MapHandler(parsedYaml), nil
}
