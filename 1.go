package main

import (
	"fmt"
	"encoding/jason"
	"log"
	"net/http"
)

func sample(rw http.ResponseWriter, t *http.Request) {
	var ht string
	
	fmt.Fprintf(w, "enter the number")
	fmt.Fscanf(r, "%s", &ht)
	fmt.Fprintf(w, "this is sample \n")
	fmt.Fprintf(w, "\n XXXX  this is mine %s", ht)
}

func main() {
	http.HandleFunc("/", sample)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
