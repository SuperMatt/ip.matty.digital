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

func userAgent(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=UTF-8")
	ua := r.UserAgent()
	fmt.Fprintf(w, ua+"\n")
}

func main() {
	http.HandleFunc("/", ipHandler)
	http.HandleFunc("/ua", userAgent)
	appengine.Main()
}
