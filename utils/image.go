package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

const (
	UploadDir = "./uploads"
	MaxSize   = 5 * 1024 * 1024 // 5MB
)

// Initialize upload directory
func InitUploadDir() error {
	if err := os.MkdirAll(UploadDir, 0755); err != nil {
		return fmt.Errorf("failed to create upload directory: %v", err)
	}
	return nil
}

// SaveImage saves an uploaded image and returns the file path
func SaveImage(file *multipart.FileHeader, prefix string) (string, error) {
	// Check file size
	if file.Size > MaxSize {
		return "", fmt.Errorf("file too large: max size is %d bytes", MaxSize)
	}

	// Check file type
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !isValidImageType(ext) {
		return "", fmt.Errorf("invalid file type: only jpg, jpeg, png, and gif are allowed")
	}

	// Generate unique filename
	filename := fmt.Sprintf("%s_%s%s", prefix, uuid.New().String(), ext)
	filepath := filepath.Join(UploadDir, filename)

	// Open source file
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open uploaded file: %v", err)
	}
	defer src.Close()

	// Create destination file
	dst, err := os.Create(filepath)
	if err != nil {
		return "", fmt.Errorf("failed to create destination file: %v", err)
	}
	defer dst.Close()

	// Copy file content
	if _, err = io.Copy(dst, src); err != nil {
		return "", fmt.Errorf("failed to save file: %v", err)
	}

	// Return relative path
	return fmt.Sprintf("/uploads/%s", filename), nil
}

// isValidImageType checks if the file extension is a valid image type
func isValidImageType(ext string) bool {
	validTypes := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
	}
	return validTypes[ext]
}

// DeleteImage deletes an image file
func DeleteImage(path string) error {
	if path == "" {
		return nil
	}

	// Remove /uploads/ prefix if present
	path = strings.TrimPrefix(path, "/uploads/")
	filepath := filepath.Join(UploadDir, path)

	// Check if file exists
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return nil
	}

	// Delete file
	if err := os.Remove(filepath); err != nil {
		return fmt.Errorf("failed to delete file: %v", err)
	}

	return nil
}
