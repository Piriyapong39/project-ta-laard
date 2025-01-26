package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func UploadPicture(productID string) error {

	baseDir := filepath.Join("..", "upload", "image")

	if err := os.MkdirAll(baseDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create base upload directory: %w", err)
	}

	today := time.Now().Format("2006-01-02")
	todayDir := filepath.Join(baseDir, today, productID)

	if err := os.MkdirAll(todayDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create daily upload directory: %w", err)
	}

	return nil
}
