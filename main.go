package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
)

func ipHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=UTF-8")
	hp := r.RemoteAddr
	fmt.Fprintf(w, hp+"\n")
}

func main() {
	http.HandleFunc("/", ipHandler)
	appengine.Main()
}
