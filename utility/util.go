package utility

import (
	"io"
	"log"
	"net/http"
	"os"
)

func DownloadPhoto(photoURL string) {
	res, err := http.Get(photoURL)
	if err != nil {
		log.Println("Error while Downloading photo")
	}
	defer res.Body.Close()
	file, err := os.Create("./resources/" + photoURL)
	if err != nil {
		log.Printf("Error while creating photo file %v\n", err)
	}
	_, err = io.Copy(file, res.Body)
	if err != nil {
		log.Printf("%v\n", err)
	}
}
