package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// temp file
	tempFile, err := os.CreateTemp("", "logs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		fmt.Println("Removing tempFile", tempFile.Name())
		if err = os.Remove(tempFile.Name()); err != nil {
			log.Fatal(err)
		}
	}()

	_, err = tempFile.Write([]byte("Hello World\n"))
	if err != nil {
		log.Fatal(err)
		tempFile.Close()
		return
	}
	defer tempFile.Close()

	// temp dir
	tempDir, err := os.MkdirTemp("", "my_app_logs")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		fmt.Println("Removing tempDir", tempDir)
		if err = os.RemoveAll(tempDir); err != nil {
			log.Fatal(err)
		}
	}()
}
