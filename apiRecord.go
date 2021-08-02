package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"
)

func apiRecord(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Println("Unsupported method.")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Unsupported method."))
		return
	}
	id := r.Header.Get("X-ID")
	// TODO zkontrolovat "id"
	f, err := os.OpenFile(path.Join("data", id), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	size, err := io.Copy(f, r.Body)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Printf("[%s] Recieved chank id: %s, size: %d\n", time.Now().String(), id, size)
	w.WriteHeader(http.StatusOK)
}
