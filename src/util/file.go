package util

import (
	"fmt"
	"path/filepath"
	"strings"

	"os"
)

var fileExtensions = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
	".bmp":  true,
	".webp": true,
	".svg":  true,
	".ico":  true,
	".tiff": true,
	".pdf":  true,
	".doc":  true,
	".docx": true,
	".xls":  true,
	".xlsx": true,
	".ppt":  true,
	".pptx": true,
	".txt":  true,
	".csv":  true,
	".rtf":  true,
	".odt":  true,
	".mp4":  true,
	".avi":  true,
	".mov":  true,
	".mkv":  true,
	".wmv":  true,
	".flv":  true,
	".webm": true,
}

func ExtractFileURL(fullURL string) string {
	prefix := "/storage/file/"
	if pos := strings.Index(fullURL, prefix); pos != -1 {
		return fullURL[pos+len(prefix):]
	}

	return fullURL
}

func isFile(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	return fileExtensions[ext]
}

func MakeFullURL(path string) string {
	if !isFile(path) {
		return path
	}

	prefix := "storage/file"
	domain := fmt.Sprintf("%s/%s", os.Getenv("BACKEND_URL"), prefix)
	if domain == "" || strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") || path == "" {
		return path
	}

	return fmt.Sprintf("%s/%s", domain, path)
}
