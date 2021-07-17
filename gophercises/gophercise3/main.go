package main

import (
	"fmt"
	"os"
)

func main() {
	var files []string = []string{"sample1.html", "sample2.html", "sample3.html", "sample4.html", }
	for _, file := range files {
		htmlData, err := os.Open(file)
		if err != nil {
			panic(err)
		}
		var links []Link
		links, err = Parse(htmlData)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%#v\n", links)
	}
}


