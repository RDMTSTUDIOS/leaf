package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"path/filepath"
)

func main() {

	currentDir, err := filepath.Abs(".")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("\n.\n")
	tf, ts := printTree(currentDir, "")
	fmt.Printf("\n\033[38;5;243mtotal %d: \033[38;5;23m", tf)

	if ts < 1024*1024 {
		fileSizeKB := math.Round(float64(ts)/1024*100) / 100
		fmt.Printf("(%.2f KB)\033[0m\n\n", fileSizeKB)
	} else {
		fileSizeMB := math.Round(float64(ts)/(1024*1024)*100) / 100
		fmt.Printf("(%.2f MB)\033[0m\n\n", fileSizeMB)
	}

}

var TotalFiles int = 0
var TotalSize int64 = 0

func printTree(dirPath string, prefix string) (int, int64) {

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Error:", err)
		return TotalFiles, TotalSize
	}

	for i, file := range files {
		TotalFiles++

		var itemPrefix string
		if i == len(files)-1 {
			itemPrefix = prefix + "└─ "
		} else {
			itemPrefix = prefix + "├─ "
		}

		fmt.Print(itemPrefix)

		if file.Mode().IsRegular() {
			fileSize := float64(file.Size())
			TotalSize += file.Size()

			var fileSizeStr string

			if fileSize < 1024*1024 {
				fileSizeKB := math.Round(fileSize/1024*100) / 100
				fileSizeStr = fmt.Sprintf("(%.2f KB)", fileSizeKB)
			} else {
				fileSizeMB := math.Round(fileSize/(1024*1024)*100) / 100
				fileSizeStr = fmt.Sprintf("(%.2f MB)", fileSizeMB)
			}

			fmt.Print("\033[38;5;253m", file.Name(), "\033[38;5;23m ", fileSizeStr, "\033[0m")
		}

		if file.Mode().IsDir() {
			fmt.Print("\033[38;5;249m", file.Name(), "/\033[0m")
		}

		fmt.Println()

		if file.Mode().IsDir() {
			subdirPath := filepath.Join(dirPath, file.Name())
			printTree(subdirPath, prefix+"│   ")
		}
	}
	return TotalFiles, TotalSize
}
