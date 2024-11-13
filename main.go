package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal("File name is required!")
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Work directory error: %s", err)
	}

	filepath := fmt.Sprintf("%s/%s", cwd, os.Args[1])

	if _, err := os.Stat(filepath); errors.Is(err, os.ErrNotExist) {
		log.Fatalf("The file %s does not exist!", os.Args[1])
	}

	p := flag.String("p", "8080", "port")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath)
	})

	if err := http.ListenAndServe(fmt.Sprintf(":%s", *p), nil); err != nil {
		log.Fatalf("Serve error: %s", err)
	}

}
