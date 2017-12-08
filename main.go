package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, KubeCon!")
}

func main() {
	http.HandleFunc("/", handler)
	port := os.Getenv("APP_PORT")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
