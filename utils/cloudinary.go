package utils

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

var cld *cloudinary.Cloudinary

func InitCloudinary() error {
	var err error
	cld, err = cloudinary.NewFromParams(
		os.Getenv("CLOUDINARY_CLOUD_NAME"),
		os.Getenv("CLOUDINARY_API_KEY"),
		os.Getenv("CLOUDINARY_API_SECRET"),
	)
	return err
}

func UploadToCloudinary(file *multipart.FileHeader, folder string) (string, error) {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Validate file
	if file == nil {
		return "", fmt.Errorf("file is nil")
	}

	if file.Size == 0 {
		return "", fmt.Errorf("file is empty")
	}

	// Check file type
	contentType := file.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") {
		return "", fmt.Errorf("unsupported file type: %s", contentType)
	}

	// Open file
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer src.Close()

	// Upload to Cloudinary
	uploadResult, err := cld.Upload.Upload(
		ctx,
		src,
		uploader.UploadParams{
			Folder:       folder,
			ResourceType: "image",
		},
	)
	if err != nil {
		return "", fmt.Errorf("cloudinary upload failed: %v", err)
	}

	if uploadResult.SecureURL == "" {
		return "", fmt.Errorf("cloudinary returned empty URL")
	}

	return uploadResult.SecureURL, nil
}

func DeleteFromCloudinary(publicID string) error {
	_, err := cld.Upload.Destroy(
		context.Background(),
		uploader.DestroyParams{
			PublicID: publicID,
		},
	)
	return err
}

// Extract public ID from Cloudinary URL
func GetPublicIDFromURL(url string) string {
	// Remove the domain and version
	parts := filepath.Base(url)
	// Remove the file extension
	return parts[:len(parts)-len(filepath.Ext(parts))]
}
