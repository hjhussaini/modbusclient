package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/hjhussaini/plc/api"
)

func main() {
	url := os.Getenv("URL")

	router := mux.NewRouter()

	plc := api.NewPLC()
	plc.RegisterAPIs(router)

	server := http.Server{
		Addr:         url,
		Handler:      router,
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Println(err)

			os.Exit(1)
		}
	}()
	fmt.Printf("PLCServer is running on %s ...\n", url)

	killSignal := make(chan os.Signal)
	signal.Notify(killSignal, os.Interrupt)
	signal.Notify(killSignal, os.Kill)

	signaled := <-killSignal
	fmt.Printf("PLCServer is shutting down (%sed)\n", signaled)

	timeout, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeout)
}
