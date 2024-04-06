package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/onlyhavecans/zipcheck/internal/ziptools"
)

func getDirectoriesFromArgs() []string {
	args := os.Args[1:]
	return args
}

func getAllZipFilesIn(directory string) ([]string, error) {
	// walk through all directories and files
	// put all zip files in a slice and return the slice
	var allZipFiles []string
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filepath.Ext(path) == ".zip" {
			allZipFiles = append(allZipFiles, path)
		}
		return nil
	})
	return allZipFiles, err
}

func checkFile(file string) {
	if ziptools.IsValidZip(file) {
		fmt.Print(".")
	} else {
		fmt.Println("Invalid zip file:", file)
	}
}

func main() {
	const maxGoroutines = 10

	// get all directories passed as arguments and put them in a slice
	directories := getDirectoriesFromArgs()
	fmt.Println("Passed Directories:", directories)

	// walk through all directories and check all zip files
	for _, directory := range directories {
		fmt.Println("Checking Directory:", directory)
		allFiles, err := getAllZipFilesIn(directory)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("Found", len(allFiles), "zip files in:", directory)

		// create a semaphore to limit the number of goroutines
		semaphore := make(chan struct{}, maxGoroutines)
		var wg sync.WaitGroup
		for _, file := range allFiles {
			wg.Add(1)
			semaphore <- struct{}{}
			go func(file string) {
				defer wg.Done()
				checkFile(file)
				<-semaphore
			}(file)
		}
		wg.Wait()
	}
}
