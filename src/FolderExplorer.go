package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func explorer(out io.Writer, path string, printfiles bool, indent string) error {
	entries, err := os.ReadDir(path) // View folder contents
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries { // Iterating over content
		fmt.Fprintln(out, indent+entry.Name()) // Content output
		if entry.IsDir() {                     // If the content is a folder, we call the function recursively, adding new indentation for more readability
			newIndent := indent + " └─ "
			if err := explorer(out, filepath.Join(path, entry.Name()), printfiles, newIndent); err != nil { // When called, we added the name of the folder to path so that it would go through it
				log.Fatal(err)
			}
		}
	}
	return nil
}

func printInstructions() { // Function for displaying instructions for running the utility
	fmt.Println("[*] Launch:\n\t1. go run FolderExplorer.go <folder name> -f\n\t2. make\n\t  ./FolderExplorer <folder name> -f\n\t3. go build FolderExplorer.go\n\t  ./FolderExplorer <folder name> -f")
	fmt.Println("[Warning] The second compilation method requires a Makefile, which is contained in the src folder of the project source code")
	fmt.Println("[*] -f - a required flag that gives permission to read the folder")
}

func main() {
	fmt.Println("FolderExplorer [Golang] v2.0 by oni-engineer\n ")
	out := os.Stdout
	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		printInstructions()
		log.Fatal("[*] Exit")
	}
	if len(os.Args) != 3 {
		fmt.Println("[E] No arguments specified")
		printInstructions()
		log.Fatal("[E] No arguments specified")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := explorer(out, path, printFiles, "") // Here you can replace the indentation character with some other one
	if err != nil {
		log.Fatal(err)
	}
}
