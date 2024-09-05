package utility

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func GetFilesInDir(dir string) {

	var fileList []string

	// Use filepath.Walk to iterate through the directory
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Only add files to the list (ignore directories)
		if !info.IsDir() {
			fileList = append(fileList, path)
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Error walking the directory: %v", err)
	}
	processFiles(fileList)

}

// Function to process the files
func processFiles(files []string) {
	fmt.Println("Processing the following files:")
	for _, file := range files {
		// Read the file content
		content, err := readFileContent(file)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", file, err)
			continue
		}

		// Print the file name and content
		fmt.Printf("File: %s\nContent:\n%s\n", file, content)
	}
}

// Function to read the content of a file
func readFileContent(filePath string) (string, error) {
	// Read the file contents
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// Convert the content to a string and return
	return string(content), nil
}
