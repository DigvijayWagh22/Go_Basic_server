package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/books/{title}/page/{pageNo}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		Title := vars["title"]
		Number := vars["pageNo"]
		fmt.Fprintf(w, "title is %s and page no is %s", Title, Number)
	})
	http.ListenAndServe(":80", r)
}
