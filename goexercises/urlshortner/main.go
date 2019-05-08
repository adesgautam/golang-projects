package main

import (
	"fmt"
	"net/http"
	"flag"
	"os"
	"io/ioutil"

	//"github.com/gophercises/urlshort"
	"./handler"
)

func main() {
	yamlFile := flag.String("yaml", "urls.yaml", "a yaml file with path and url defined")
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string {
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)
	fmt.Println(mapHandler)
	// Build the YAMLHandler using the mapHandler as the fallback
	
/*	
	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
*/  
	yaml, err := ioutil.ReadFile(*yamlFile)
	if err != nil {
		fmt.Println("Failed to read file 'urls.yaml'")
		os.Exit(1)
	}

	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
	
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
















