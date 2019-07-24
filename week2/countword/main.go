package main

import (
	"fmt"
	"grab/week2/countword/file"
	"log"
)

const (
	dataRoot = "data"
)

func printFileData(file chan file.FileData) {
	for data := range file {
		fmt.Println(data.Words)
	}
}

func main() {
	in, err := file.ReadTextFiles(dataRoot)
	if err != nil {
		log.Fatal("Failed to read files: ", err)
	}
	go printFileData(in)
	fmt.Scanln()
}
