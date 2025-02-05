package main

//http.HandleFunc
import (
	"fmt"
	"log"
	"net/http"
)

func goodHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "oh nice to meet you felani %s", r.URL.Path[len("/good/"):])
}
func badHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "go out %s", r.URL.Path[len("/bad/"):])
}
func main() {
	http.HandleFunc("/good/", goodHandler)
	http.HandleFunc("/bad/", badHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
