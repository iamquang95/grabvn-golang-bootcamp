package main

import (
	"fmt"
	"grab/week2/countword/fileutils"
	"log"
	"sync"
	"time"
)

const (
	dataRoot = "data"
)

type result struct {
	wordsCount map[string]int
}

func countWords(file fileutils.FileData) (result, error) {
	cnt := make(map[string]int)
	if file.Err != nil {
		return result{cnt}, file.Err
	}
	for _, word := range file.Words {
		cnt[word]++
	}
	return result{cnt}, nil
}

func countWordsInFiles(files chan fileutils.FileData, counts chan result, wg *sync.WaitGroup) {
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

func collectResults(counts chan result, wg *sync.WaitGroup) {
	res := make(map[string]int)
	for count := range counts {
		for k, v := range count.wordsCount {
			res[k] = res[k] + v
		}
	}
	fmt.Println(res)
	wg.Done()
}

func pool(files chan fileutils.FileData, counts chan result, nFiles int, wg *sync.WaitGroup) {
	for i := 0; i < nFiles; i++ {
		go countWordsInFiles(files, counts, wg)
	}
}

func main() {
	timeStart := time.Now()

	filePaths, err := fileutils.ListAllFiles(dataRoot)
	if err != nil {
		log.Fatal("Failed to read files under ", dataRoot, ": ", err)
	}

	// Add a wait group that equal number of files in folder
	var wgFiles sync.WaitGroup
	wgFiles.Add(len(filePaths))

	var wgCollectResult sync.WaitGroup
	wgCollectResult.Add(1)

	files := fileutils.ReadTextFiles(filePaths)
	counts := make(chan result)

	go pool(files, counts, len(filePaths), &wgFiles)
	// go countWordsInFiles(files, counts, &wgFiles)
	go collectResults(counts, &wgCollectResult)
	go func() {
		wgFiles.Wait()
		close(counts)
	}()
	wgCollectResult.Wait()
	timeEnd := time.Now()
	log.Output(2, fmt.Sprintf("wordCounter finished in %s", timeEnd.Sub(timeStart)))
}
