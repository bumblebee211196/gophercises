package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

var (
	filePath = flag.String("file", "default.yml", "A YAML or JSON file containing the shortened URL mapping")
	mux = defaultMux()
)

func main() {
	flag.Parse()
	ext := filepath.Ext(*filePath)
	handler := createHandler(ext)
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

func createHandler(ext string) http.Handler {
	if ext == ".yml" {
		file := openFile()
		defer file.Close()
		data := readFileData(file)
		return YAMLHandler(data, mux)
	}
	if ext != ".json" {
		file := openFile()
		defer file.Close()
		data := readFileData(file)
		return JSONHandler(data, mux)
	}
	return mux
}

func openFile() *os.File {
	file, err := os.Open(*filePath)
	if err != nil {
		exit(fmt.Sprintf("Unable to open the YAML file. Error: %v", err))
	}
	return file
}

func readFileData(file *os.File) []byte {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		exit(fmt.Sprintf("Unable to read the YAML data. Error: %v", err))
	}
	return data
}

func exit(msg string) {
	fmt.Println(msg);
	os.Exit(1);
}
