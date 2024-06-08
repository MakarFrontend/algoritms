package main

import (
	"fmt"
	"os"
)

func walkDir(path string) {
	f, er := os.Open(path)
	if er != nil {
		fmt.Println(er)
		return
	}
	files, err := f.ReadDir(-1)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Printf("Folder: %s\n", file.Name())
			walkDir(path + "/" + file.Name())
		} else {
			fmt.Printf("File: %s\n", file.Name())
		}
	}
}

func main() {
	walkDir(".")
}
