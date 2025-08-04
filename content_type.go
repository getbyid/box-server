package main

import (
	"strings"
)

// https://developer.mozilla.org/en-US/docs/Web/HTTP/Guides/MIME_types/Common_types
var textTypes = map[string]string{
	".css":   "text/css",
	".csv":   "text/csv",
	".htm":   "text/html",
	".html":  "text/html",
	".js":    "text/javascript",
	".json":  "application/json",
	".md":    "text/markdown",
	".mjs":   "text/javascript",
	".svg":   "image/svg+xml",
	".txt":   "text/plain",
	".xhtml": "application/xhtml+xml",
	".xml":   "application/xml",
}

// textContentType определяет тип контента по расширению файла.
// Это имеет смысл только для текстовых файлов, т.к. бинарных форматов
// слишком много и их можно определить другим способом.
func textContentType(filename string) string {
	for k, v := range textTypes {
		if strings.HasSuffix(filename, k) || strings.Contains(filename, k+"?") {
			return v
		}
	}
	return ""
}
