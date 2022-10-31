package main

import (
	"log"
	"net/http"
)

func send(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5000")
	w.Header().Set("Access-Control-Max-Age", "86400")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	c := http.Cookie{Name: "some", Value: "cookie"}
	http.SetCookie(w, &c)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/send", send)
	mux.HandleFunc("/send/test", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r)
		w.Write([]byte("hello"))
	})
	http.ListenAndServe("127.0.0.1:8080", mux)
}
