package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"practice/Go训练营/3.第二周作业/api"
	"syscall"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/user", api.GetUserById)
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
			}
		}()
		server.ListenAndServe()
	}()

	q := make(chan os.Signal)
	signal.Notify(q, syscall.SIGINT, syscall.SIGTERM)
	<-q
}
