package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	httpMux := http.NewServeMux()
	httpMux.HandleFunc("/record", apiRecord)

	s := &http.Server{
		Addr:           ":http",
		Handler:        httpMux,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
