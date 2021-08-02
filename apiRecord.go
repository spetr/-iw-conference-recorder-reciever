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
		fmt.Printf("[%s] Unsupported request method\n", time.Now().String())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Unsupported request method."))
		return
	}
	id := r.Header.Get("X-ID")
	sequenceID := r.Header.Get("X-SeqID")
	endParam := r.URL.Query()["end"]
	// TODO zkontrolovat "id" a "sequenceID"
	f, err := os.OpenFile(path.Join("data", id), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Printf("[%s] Error in file close. ID: %s, SeqID: %s, Error: %s\n", time.Now().String(), id, sequenceID, err.Error())
		}
	}()
	size, err := io.Copy(f, r.Body)
	if err != nil {
		fmt.Printf("[%s] Error in file write. ID: %s, SeqID: %s, Error: %s\n", time.Now().String(), id, sequenceID, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Printf("[%s] Recieved chunk. ID: %s, SeqID: %s, Size: %d\n", time.Now().String(), id, sequenceID, size)
	w.WriteHeader(http.StatusOK)

	if len(endParam) > 0 {
		fmt.Printf("[%s] Finished job. ID: %s\n", time.Now().String(), id)
	}
}
