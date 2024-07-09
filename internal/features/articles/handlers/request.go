package handlers

import (
	"context"
	"io"
	"newsapps/internal/features/articles"
	"os"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

// type LoginRequest struct {
// 	Email    string `json:"email" form:"email"`
// 	Password string `json:"password" form:"password"`
// }

type NewArticlesRequest struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
}

// type AlamatRequest struct {
// 	Alamat string `json:"alamat"`
// }

// func ToModelUsers(r RegisterRequest) users.User {
// 	return users.User{
// 		Name:     r.Name,
// 		Password: r.Password,
// 		Email:    r.Email,
// 		Phone:    r.Phone,
// 	}
// }

func ToArticleEntity(a NewArticlesRequest) articles.Article {
	return articles.Article{

		Title:   a.Title,
		Content: a.Content,
	}
}

func (a NewArticlesRequest) uploadToCloudinary(file io.Reader, filename string) (string, error) {
	// Konfigurasi Cloudinary

	cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	if err != nil {
		return "", err
	}

	// Upload file ke Cloudinary
	uploadParams := uploader.UploadParams{
		Folder:   "articles_pictures",
		PublicID: filename,
	}
	uploadResult, err := cld.Upload.Upload(context.Background(), file, uploadParams)
	if err != nil {
		return "", err
	}

	// Ambil URL publik dari hasil unggah
	publicURL := uploadResult.URL
	return publicURL, nil
}
