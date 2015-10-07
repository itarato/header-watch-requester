package main

import (
	"log"
	"net/http"
	"time"
)

type ServerHandler struct {
}

func (sh *ServerHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	log.Println("called http.ServeHTTP")
}

func main() {
	log.Println("crawler has been started")

	server := &http.Server{
		Addr:           "localhost:8080",
		Handler:        &ServerHandler{},
		ReadTimeout:    time.Second * 60,
		WriteTimeout:   time.Second * 60,
		MaxHeaderBytes: 0,
	}

	log.Println(server)

	err := server.ListenAndServe()
	if err != nil {
		log.Panicln(err)
	}

	for {
	}
}
