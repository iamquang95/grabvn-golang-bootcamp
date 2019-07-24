package file

import (
	"bufio"
	"os"
	"path/filepath"
)

// FileData describe a text file with path, words in file and error if any
type FileData struct {
	Path  string
	Words []string
	Err   error
}

// ReadTextFiles for each file under root, read words and put words to channel words
func ReadTextFiles(root string) (chan FileData, error) {
	files := make(chan FileData)
	paths, err := listAllFiles(root)
	if err != nil {
		return files, err
	}
	go func() {
		for path := range paths {
			words, err := extractFileToWords(path)
			files <- FileData{path, words, err}
		}
	}()
	return files, err
}

func extractFileToWords(path string) (words []string, err error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	// fmt.Println("Got ", words, " in file ", path)
	return
}

// listAllFiles list all file paths under a root
func listAllFiles(root string) (chan string, error) {
	paths := make(chan string)
	var err error
	go func() {
		defer close(paths)
		err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			// fmt.Println("Got file ", path)
			paths <- path
			return nil
		})
	}()
	return paths, err
}
