package main

import (
	"fmt"
	"grab/week2/countword/file"
	"log"
	"sync"
)

const (
	dataRoot = "data"
)

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

func countWordsInFiles(files chan file.FileData, counts chan result, wg *sync.WaitGroup) {
	for file := range files {
		if file.Err != nil {
			log.Fatal("Failed to read file", file.Path)
			return
		}
		wordsCnt, _ := countWords(file)
		counts <- wordsCnt
		wg.Done()
	}
}

func printResult(res result) {
	for k, v := range res.wordsCount {
		fmt.Println(k, ":", v)
	}
}

func collectResults(counts chan result, wg *sync.WaitGroup) {
	res := make(map[string]int)
	for count := range counts {
		for k, v := range count.wordsCount {
			res[k] = res[k] + v
		}
	}
	printResult(result{res})
	wg.Done()
}

func main() {
	filePaths, err := file.ListAllFiles(dataRoot)
	if err != nil {
		log.Fatal("Failed to read files under ", dataRoot, ": ", err)
	}

	// Add a wait group that equal number of files in folder
	var wgFiles sync.WaitGroup
	wgFiles.Add(len(filePaths))

	var wgCollectResult sync.WaitGroup
	wgCollectResult.Add(1)

	files := file.ReadTextFiles(filePaths)
	counts := make(chan result)

	go countWordsInFiles(files, counts, &wgFiles)
	go collectResults(counts, &wgCollectResult)
	go func() {
		wgFiles.Wait()
		close(counts)
	}()
	wgCollectResult.Wait()
}
