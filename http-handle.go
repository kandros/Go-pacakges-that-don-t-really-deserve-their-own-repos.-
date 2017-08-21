package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	addr := ":3000"
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Ciaone!")
}
