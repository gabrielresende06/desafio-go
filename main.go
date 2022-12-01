package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("test"))
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running ...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
