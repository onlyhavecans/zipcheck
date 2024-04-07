package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"

	"github.com/onlyhavecans/zipcheck/internal/ziptools"
)

const (
	maxGoroutines = 10
	exitFail      = 1
	fileExtension = ".zip"
)

func main() {
	if err := run(os.Args, os.Stdout, os.Stderr); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(exitFail)
	}
}

func run(args []string, stdout io.Writer, stderr io.Writer) error {
	// get all directories passed as arguments and put them in a slice
	directories := getDirectoriesFromArgs(args)
	log(stdout, "Passed Directories:", directories)

	// walk through all directories and check all zip files
	for _, directory := range directories {
		allFiles, err := getFilesRecursive(directory, fileExtension)
		if err != nil {
			log(stderr, err)
			continue
		}

		log(stdout, "Found", len(allFiles), "zip files in:", directory)

		// create a semaphore to limit the number of goroutines
		semaphore := make(chan struct{}, maxGoroutines)
		var wg sync.WaitGroup
		for _, file := range allFiles {
			wg.Add(1)
			semaphore <- struct{}{}
			go func(file string) {
				defer wg.Done()
				if ziptools.IsValidZip(file) {
					log(stdout, ".")
				} else {
					log(stderr, "Invalid zip file:", file)
				}
				<-semaphore
			}(file)
		}
		wg.Wait()
	}
	return nil
}

func log(to io.Writer, v ...interface{}) {
	fmt.Fprintln(to, v...)
}

func getDirectoriesFromArgs(args []string) []string {
	return args[1:]
}

func getFilesRecursive(directory string, extension string) ([]string, error) {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		return nil, fmt.Errorf("directory %s does not exist", directory)
	}

	var allZipFiles []string
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filepath.Ext(path) == extension {
			allZipFiles = append(allZipFiles, path)
		}
		return nil
	})
	return allZipFiles, err
}
