package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func apiRecord(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Println("Unsupported method.")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Unsupported method."))
		return
	}
	id := r.Header.Get("X-ID")
	f, err := os.Create(path.Join("data", id))
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = io.Copy(f, r.Body)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
