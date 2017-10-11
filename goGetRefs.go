package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {

	var downloads []string
	downloads = os.Args[1:]

	var urls []string
	for _, url := range downloads {
		paths := strings.Split(url, ",")
		for _, each := range paths {
			urls = append(urls, each)
		}
	}

	for _, url := range urls {
		downloadFile(url)
	}
}

// -------------------------------- //

func downloadFile(url string) {
	cmd := exec.Command("wget", url)
	log.Println("Beginning Download of: ", url)
	downloadErr := cmd.Run()
	if downloadErr != nil {
		fmt.Println(downloadErr)
	}
}

// -------------------------------- //
