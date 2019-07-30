package main

import (
	"bufio"
	"fmt"
	"grab/week2/countwordworker/workerpool"
	"log"
	"os"
	"path/filepath"
)

const (
	dataDir  = "data"
	nWorkers = 1
)

func getAllFilePaths(root string) (paths []string, err error) {
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}
		// fmt.Println("Got file ", path)
		paths = append(paths, path)
		return nil
	})
	return paths, err
}

type FileJob struct {
	path           string
	wordCntChannel chan map[string]int
}

func (fileJob FileJob) Run() {
	file, err := os.Open(fileJob.path)
	defer file.Close()
	if err != nil {
		// Still need to return to channel an empty map, since collect result need n items in channel
		fileJob.wordCntChannel <- make(map[string]int)
		fmt.Println("failed to read file", fileJob.path, "with err =", err)
		return
	}

	cntWords := countWordsInFile(file)
	fileJob.wordCntChannel <- cntWords
	return
}

func countWordsInFile(file *os.File) map[string]int {
	cnt := make(map[string]int)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		cnt[word]++
	}
	return cnt
}

func collectResult(wordCntChannel <-chan map[string]int, nFiles int) {
	res := make(map[string]int)
	for i := 0; i < nFiles; i++ {
		wordCnt := <-wordCntChannel
		for k, v := range wordCnt {
			res[k] = res[k] + v
		}
	}
	fmt.Println(res)
}

func main() {
	filepaths, err := getAllFilePaths(dataDir)
	if err != nil {
		log.Fatal("failed to read all file path", err)
	}

	p := workerpool.NewWorkerPool(nWorkers)
	resultChannel := make(chan map[string]int)

	for _, filepath := range filepaths {
		p.Dispatch(FileJob{filepath, resultChannel})
	}
	collectResult(resultChannel, len(filepaths))
}
