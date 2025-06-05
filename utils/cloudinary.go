package utils

import (
	"context"
	"mime/multipart"
	"os"
	"path/filepath"

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
	// Open file
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Upload to Cloudinary
	uploadResult, err := cld.Upload.Upload(
		context.Background(),
		src,
		uploader.UploadParams{
			Folder:       folder,
			ResourceType: "image",
		},
	)
	if err != nil {
		return "", err
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
