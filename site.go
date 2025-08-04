package main

import (
	"archive/zip"
	"fmt"
	"io"
	"strings"
)

type Site struct {
	index string
	files map[string]*zip.File
}

func NewSite(index string) *Site {
	return &Site{index, make(map[string]*zip.File)}
}

// findPrefix просматривает все пути внутри архива
// и пытается найти единственную вложенную папку.
func findPrefix(r *zip.ReadCloser) string {
	var prefix string
	for i, f := range r.File {
		if i == 0 {
			// у первого файла берём папку
			before, _, found := strings.Cut(f.Name, "/")
			if found {
				prefix = before + "/"
			}
		} else {
			// у остальных проверяем её наличие
			if !strings.HasPrefix(f.Name, prefix) {
				prefix = ""
				break
			}
		}
	}
	return prefix
}

// LoadFromZip заполняет карту сайта, для каждого пути к файлу
// сохраняет дескриптор для чтения его содержимого.
func (s *Site) LoadFromZip(filename string) error {
	r, err := zip.OpenReader(filename)
	if err != nil {
		return fmt.Errorf("zip reader: %w", err)
	}
	// defer r.Close()

	prefix := findPrefix(r)
	for _, f := range r.File {
		s.files[strings.TrimPrefix(f.Name, prefix)] = f
	}

	return nil
}

// ContentByPath возвращает интерфейс для получения контента
func (s *Site) ContentByPath(path string) (io.ReadCloser, error) {
	f, ok := s.files[path[1:]]
	if !ok {
		return nil, fmt.Errorf("not found: %s", path)
	}

	return f.Open()
}
