package main

import (
	"fmt"
	"log"
	"net/http"
)

func wakeUp(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" || r.Method != "GET" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	_, err := w.Write([]byte("OK"))
	if err != nil {
		log.Default().Println(err)
	}
}

func Caffeine() {
	http.HandleFunc("/", wakeUp)
	addr := ":8080"
	if Port != "" {
		addr = fmt.Sprintf(":%s", Port)
	}
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
