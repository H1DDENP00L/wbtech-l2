package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// downloadFile загружает файл из указанного URL и сохраняет его по указанному пути
func downloadFile(url, outputPath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("не удалось выполнить запрос: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("сервер вернул ошибку: %v", resp.Status)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("не удалось создать файл: %v", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("ошибка при записи данных: %v", err)
	}

	return nil
}

// downloadWebsite загружает все файлы сайта по указанному URL
func downloadWebsite(baseURL, outputDir string) error {
	if !strings.HasSuffix(baseURL, "/") {
		baseURL += "/"
	}

	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return fmt.Errorf("не удалось создать директорию: %v", err)
	}

	indexPath := filepath.Join(outputDir, "index.html")
	fmt.Printf("Загрузка %s в %s\n", baseURL, indexPath)

	err := downloadFile(baseURL, indexPath)
	if err != nil {
		return fmt.Errorf("ошибка загрузки главной страницы: %v", err)
	}

	fmt.Println("Загрузка завершена.")
	return nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Использование: wget [URL] [путь_для_сохранения]")
		os.Exit(1)
	}

	url := os.Args[1]
	outputDir := os.Args[2]

	if err := downloadWebsite(url, outputDir); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err)
		os.Exit(1)
	}
}
