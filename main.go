package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
)

func ipHandler(w http.ResponseWriter, r *http.Request) {
	hp := r.RemoteAddr
	fmt.Fprintf(w, hp)
}

func main() {
	http.HandleFunc("/", ipHandler)
	appengine.Main()
}
