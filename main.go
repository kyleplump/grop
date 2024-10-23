package main

import (
	"bufio"
	"flag"
	"fmt"
	color "github.com/fatih/color"
	"io/fs"
	"os"
	"strings"
	"sync"
)

var group sync.WaitGroup

func main() {
	var directory = flag.String("d", ".", "directory to search")
	var query = flag.String("q", "", "search query")
	flag.Parse()

	root := os.DirFS(*directory)
	var paths []string

	fs.WalkDir(root, ".", func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			// todo
		}

		if path != "." {
			paths = append(paths, *directory+"/"+path)
		}

		return nil
	})

	for i := 0; i < len(paths); i++ {
		group.Add(1)
		go search(paths[i], *query)
	}
	group.Wait()
}

func search(path string, q string) {
	f, err := os.Open(path)
	defer f.Close()

	if err != nil {
		fmt.Println("error opening file: ", path)
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), q) {
			filename := "=> " + path;
			coloredFilename := color.New(color.FgGreen);
			coloredFilename.Println(filename);
			haystack := scanner.Text()
			before, after, _ := strings.Cut(haystack, q)

			fmt.Print(before)
			coloredPart := color.New(color.FgRed)
			coloredPart.Print(q)

			remainder := after

			for strings.Contains(remainder, q) {
				before, after, _ := strings.Cut(remainder, q)

				fmt.Print(before)
				coloredPart := color.New(color.FgRed)
				coloredPart.Print(q)
				fmt.Println(after)

				remainder = after
			}

			fmt.Println("");
		}
	}

	group.Done()
}
