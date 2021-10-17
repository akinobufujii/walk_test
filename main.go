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

	findFilesWalk, err := findFilesWithWalk(SearchPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	findFilesWalkDir, err := findFilesWithWalkDir(SearchPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 結果一致確認
	// まずは検出数が同じか見る
	fmt.Printf("findFilesWithWalk %v\n", len(findFilesWalk))
	fmt.Printf("findFilesWithWalkDir %v\n", len(findFilesWalkDir))

	// 次に内容一致確認
	// 検出数が同じ状態で片方をmapのキーにして、その後キーをもう片方の要素で削除したとき
	// 0になれば一致したと言えるはず
	findMap := map[string]struct{}{}
	for _, filename := range findFilesWalk {
		findMap[filename] = struct{}{}
	}

	fmt.Printf("findMap %v\n", len(findMap))

	for _, filename := range findFilesWalkDir {
		delete(findMap, filename)
	}

	fmt.Printf("findMap %v\n", len(findMap))
}
