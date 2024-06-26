package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please provide an argument")
		os.Exit(1)
	}
	file := args[1]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)
	// fmt.Println("Pathsplit: ", pathSplit)

	for _, dir := range pathSplit {
		// fmt.Println("Dir: ", dir)
		fullPath := filepath.Join(dir, file)
		// fmt.Println("Fullpath: ", fullPath)

		// Does it exist?
		fileInfo, err := os.Stat(fullPath)
		if err != nil {
			continue
		}

		mode := fileInfo.Mode()
		fmt.Println("Mode: ", mode)
		if !mode.IsRegular() {
			continue
		}

		// Is it executable?
		if mode&0111 != 0 {
			fmt.Println("found:", fullPath)
			os.Exit(0)
		}
	}
}
