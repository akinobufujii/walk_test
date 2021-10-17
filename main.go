package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func findFilesWithWalk(root string) ([]string, error) {
	findList := []string{}

	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return errors.Wrap(err, "failed filepath.Walk")
		}

		if info.IsDir() {
			return nil
		}

		findList = append(findList, path)
		return nil
	})
	return findList, err
}

func findFilesWithWalkDir(root string) ([]string, error) {
	findList := []string{}

	err := filepath.WalkDir(root, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return errors.Wrap(err, "failed filepath.Walk")
		}

		if info.IsDir() {
			return nil
		}

		findList = append(findList, path)
		return nil
	})
	return findList, err
}

func main() {
	SearchPath := os.ExpandEnv("${GOPATH}")

	findFiles, err := findFilesWithWalk(SearchPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("findFilesWithWalk %v\n", len(findFiles))

	findFiles, err = findFilesWithWalkDir(SearchPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("findFilesWithWalkDir %v\n", len(findFiles))
}
