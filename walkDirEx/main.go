package main

import (
	"io/fs"
	"path/filepath"
)

func walk(s string, d fs.DirEntry, err error) error {
	if err != nil {
		return err //as long as there are new files keep searching else return err
	}
	if !d.IsDir() {
		println(s) //keep searching
	}
	return nil
}

func main() {
	filepath.WalkDir("..", walk)
}
