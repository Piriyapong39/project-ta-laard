package utils

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func UploadPicture(mainImage *multipart.FileHeader, c *fiber.Ctx, productID string) (string, error) {

	mainImageSplit := strings.Split(mainImage.Filename, ".")
	extName := mainImageSplit[len(mainImageSplit)-1]
	if extName != "jpg" && extName != "png" && extName != "jpeg" {
		return "", errors.New("only jpg, png, jpeg are allow")
	}
	baseDir := filepath.Join("..", "upload", "image")
	if err := os.MkdirAll(baseDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create base upload directory: %w", err)
	}
	today := time.Now().Format("2006-01-02")
	todayDir := filepath.Join(baseDir, today, productID)
	if err := os.MkdirAll(todayDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create daily upload directory: %w", err)
	}
	picPath := todayDir + "/" + mainImage.Filename
	err := c.SaveFile(mainImage, picPath)
	if err != nil {
		return "", err
	}

	return picPath, nil
}
