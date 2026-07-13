package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	filePath := "./10section/1-files/text.txt"
	data := "Welcome to the Go programming language! Lots of love for Go"
	err := os.WriteFile(filePath, []byte(data), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done writing")
	content, err := os.ReadFile(filePath) // content type is slice of bytes
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))

	/*file2, err := os.Create("file-via-create.txt") // return the file handler
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := file2.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	//defer file2.Close() // the simple way to close the file

	_, err = file2.WriteString("Welcome all Java and Python and Javascript and PHP and Ruby developers")
	if err != nil {
		log.Fatal(err)
	}*/

	fileName := "file-via-create.txt"

	printContent(fileName)

	// os.OpenFile not os.Open accepts the flags
	newFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	_, _ = newFile.WriteString(fmt.Sprintf("- C\n"))
	_, _ = newFile.WriteString(fmt.Sprintf("- Ruby\n"))
	_, _ = newFile.WriteString(fmt.Sprintf("- Fortran\n"))
	_, _ = newFile.WriteString(fmt.Sprintf("- Ada\n"))
	_, _ = newFile.WriteString(fmt.Sprintf("- Rust\n"))
}

func printContent(fileName string) {
	// when you use Open, you just use Open to read
	newFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	scanner := bufio.NewScanner(newFile) // ?the difference of .Read
	lineNum := 1
	for scanner.Scan() {
		fmt.Println(lineNum, scanner.Text())
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		if err != io.EOF { // omit the end of file error
			log.Fatal(err)
		}
	}
}
