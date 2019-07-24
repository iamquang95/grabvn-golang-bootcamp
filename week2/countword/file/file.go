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
func ReadTextFiles(paths []string) chan FileData {
	files := make(chan FileData)
	for _, path := range paths {
		go func(path string) {
			words, err := extractFileToWords(path)
			files <- FileData{path, words, err}
		}(path)
	}
	return files
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

// ListAllFiles list all file paths under a root
func ListAllFiles(root string) (paths []string, err error) {
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
