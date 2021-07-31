package main

import (
	"os"
	"path/filepath"
)

func init() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	os.Chdir(dir)
	os.Mkdir("data", 0777)
}
