package main

import (
	"fmt"
	"log"
	"net/http"
)

func pingHAndler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "pong")
}


func main() {
	fmt.Println("Started server")
	http.HandleFunc("/", pingHAndler)
    log.Fatal(http.ListenAndServe(":8000", nil))
}
