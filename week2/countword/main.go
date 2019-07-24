package main

import (
	"fmt"
	"grab/week2/countword/file"
	"log"
)

const (
	dataRoot = "data"
)

func printFileData(files chan file.FileData) {
	for data := range files {
		fmt.Println(data.Words)
	}
}

type result struct {
	wordsCount map[string]int
}

func countWords(file file.FileData) (result, error) {
	cnt := make(map[string]int)
	if file.Err != nil {
		return result{cnt}, file.Err
	}
	for _, word := range file.Words {
		cnt[word]++
	}
	return result{cnt}, nil
}

func countWordsInFiles(files chan file.FileData, counts chan result) {
	for data := range files {
		// TODO: Handle error here
		wordsCnt, _ := countWords(data)
		counts <- wordsCnt
	}
}

func collectResults(counts chan result) {
	res := make(map[string]int)
	for count := range counts {
		for k, v := range count.wordsCount {
			res[k] = res[k] + v
		}
		fmt.Println(res)
	}
}

func main() {
	in, err := file.ReadTextFiles(dataRoot)
	if err != nil {
		log.Fatal("Failed to read files: ", err)
	}
	// go printFileData(in)
	counts := make(chan result)
	go countWordsInFiles(in, counts)
	go collectResults(counts)
	fmt.Scanln()
}
