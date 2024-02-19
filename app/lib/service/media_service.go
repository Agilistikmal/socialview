package service

import (
	"io"
	"log"
	"net/http"
	"os"
)

func SaveVideo(source string, filename string) error {
	res, err := http.Get(source)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = io.Copy(f, res.Body)
	return err
}

func GetVideo(filename string) *os.File {
	f, _ := os.Open(filename)
	return f
}

func DeleteVideo(filename string) error {
	err := os.Remove(filename)
	return err
}
