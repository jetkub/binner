package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func promptUser() rune {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Do you want to replace the file? (Y/N) ")
		resp, _, err := reader.ReadRune()
		if err != nil {
			log.Fatal(err)
		}

		switch resp {
		case 'Y', 'y':
			return 'Y'
		case 'N', 'n':
			return 'N'
		default:
			fmt.Println("Invalid option. Please type Y for Yes, or N for No.")
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file to be moved to /usr/local/bin.")
	}

	srcPath := os.Args[1]
	dstPath := "/usr/local/bin/" + filepath.Base(srcPath)

	// Check if destination file exists
	if _, err := os.Stat(dstPath); err == nil {
		switch promptUser() {
		case 'Y':
			// continue to the copy and permissions part
		case 'N':
			fmt.Println("Operation cancelled by user.")
			return
		}
	}

	// Open original file
	originalFile, err := os.Open(srcPath)
	if err != nil {
		log.Fatal(err)
	}
	defer originalFile.Close()

	// Create new file
	newFile, err := os.Create(dstPath)
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	// Copy the bytes to destination from source
	_, err = io.Copy(newFile, originalFile)
	if err != nil {
		log.Fatal(err)
	}

	// Commit the file contents
	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
	}

	// Change permissions of the file to make it executable
	err = os.Chmod(dstPath, 0755)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File %s has been copied to /usr/local/bin/ and made executable.\n", srcPath)
}

