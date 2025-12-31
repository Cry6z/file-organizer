package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	targetDir := "."

	if len(os.Args) > 1 {
		targetDir = os.Args[1]
	}

	fmt.Println("Organizing folder:", targetDir)

	files, err := os.ReadDir(targetDir)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		ext := strings.ToLower(filepath.Ext(file.Name()))
		category := getCategory(ext)

		destDir := filepath.Join(targetDir, category)
		os.MkdirAll(destDir, os.ModePerm)

		srcPath := filepath.Join(targetDir, file.Name())
		destPath := filepath.Join(destDir, file.Name())

		err := os.Rename(srcPath, destPath)
		if err != nil {
			fmt.Println("Gagal memindahkan:", file.Name())
		} else {
			fmt.Println("✅", file.Name(), "→", category)
		}
	}

	fmt.Println("Selesai!")
}

func getCategory(ext string) string {
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif":
		return "images"
	case ".mp4", ".mkv", ".avi":
		return "videos"
	case ".mp3", ".wav":
		return "music"
	case ".pdf", ".docx", ".txt":
		return "docs"
	case ".zip", ".rar", ".7z":
		return "archive"
	default:
		return "others"
	}
}
