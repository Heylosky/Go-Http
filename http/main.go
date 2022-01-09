package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/healthz", healthz) //handler，谁来处理request
	http.HandleFunc("/", handler)
	http.HandleFunc("/welcome", handlerGet)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok")
}

func handler(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]string, 10)
	w.Header().Set("enVar", os.Getenv("GOPATH"))
	for k, v := range r.Header {
		//w.Header().Add(k, fmt.Sprintf("%s", v))
		m[k] = v[0]
		io.WriteString(w, fmt.Sprintf("%s: %s\n", k, v))
	}
	for k, v := range m {
		fmt.Printf("%s: %s\n", k, v)
		w.Header().Add(k, v)
	}
}

func handlerGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to <go> http server"))
	fmt.Println(r.Header["Host"])
	fmt.Println()
}
