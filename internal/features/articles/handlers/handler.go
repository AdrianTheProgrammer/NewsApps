package handlers

import (
	"newsapps/internal/features/articles"
	"newsapps/internal/helpers"
	"newsapps/internal/utils"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type ArticlesHand struct {
	srv articles.AServices
}

func NewArticlesHand(s articles.AServices) articles.AHandlers {
	return &ArticlesHand{
		srv: s,
	}
}

func (ah *ArticlesHand) ShowAllArticles() echo.HandlerFunc {
	return func(c echo.Context) error {
		responseData, err := ah.srv.ShowAllArticles()
		if err != nil {
			errCode := 500
			if strings.ContainsAny(err.Error(), "tidak ditemukan") {
				errCode = 400
			}
			return c.JSON(errCode, helpers.ResponseFormat(errCode, err.Error(), nil))
		}
		return c.JSON(200, helpers.ResponseFormat(200, "data articles berhasil diambil", responseData))
	}
}

func (ah *ArticlesHand) CreateArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		var dataUser = utils.NewTokenUtil().DecodeToken(c.Get("user").(*jwt.Token))
		var newArticle NewArticlesRequest
		err := c.Bind(&newArticle)
		if err != nil {
			return c.JSON(500, helpers.ResponseFormat(500, "article data Error", nil))
		}

		file, err := c.FormFile("articles_picture")
		var imageURL string
		if err == nil {
			// Buka file
			src, err := file.Open()
			if err != nil {
				return c.JSON(500, helpers.ResponseFormat(500, "article pic Error", nil))
			}
			defer src.Close()

			// Upload file ke Cloudinary
			imageURL, err = newArticle.uploadToCloudinary(src, file.Filename)
			if err != nil {
				return c.JSON(500, helpers.ResponseFormat(500, "article pic upload error", nil))
			}
		}
		err = ah.srv.CreateArticle(ToArticleEntity(newArticle), imageURL, dataUser.ID)
		if err != nil {
			return c.JSON(500, helpers.ResponseFormat(500, "article post error", nil))
		}
		return nil
	}

}

func (ah *ArticlesHand) DeleteArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		var dataUser = utils.NewTokenUtil().DecodeToken(c.Get("user").(*jwt.Token))
		articleID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(400, helpers.ResponseFormat(400, "article id Error", nil))
		}

		err = ah.srv.DeleteArticle(uint(articleID), dataUser.ID)
		if err != nil {
			errCode := 500
			if strings.ContainsAny(err.Error(), "tidak ditemukan") {
				errCode = 400
			}
			return c.JSON(errCode, helpers.ResponseFormat(errCode, err.Error(), nil))
		}
		return nil
	}
}

func (ah *ArticlesHand) ReadArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		articleID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(400, helpers.ResponseFormat(400, "article id Error", nil))
		}

		responseData, err := ah.srv.ReadArticle(uint(articleID))
		if err != nil {
			errCode := 500
			if strings.ContainsAny(err.Error(), "tidak ditemukan") {
				errCode = 400
			}
			return c.JSON(errCode, helpers.ResponseFormat(errCode, err.Error(), nil))
		}
		return c.JSON(200, helpers.ResponseFormat(200, "data article berhasil diambil", responseData))
	}
}

func (ah *ArticlesHand) UpdateArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		var dataUser = utils.NewTokenUtil().DecodeToken(c.Get("user").(*jwt.Token))
		articleID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(400, helpers.ResponseFormat(400, "article id Error", nil))
		}
		var newArticle NewArticlesRequest
		err = c.Bind(&newArticle)
		if err != nil {
			return c.JSON(400, helpers.ResponseFormat(400, "article data Error", nil))
		}

		file, err := c.FormFile("articles_picture")
		var imageURL string
		if err == nil {
			// Buka file
			src, err := file.Open()
			if err != nil {
				return c.JSON(500, helpers.ResponseFormat(500, "article pic Error", nil))
			}
			defer src.Close()

			// Upload file ke Cloudinary
			imageURL, err = newArticle.uploadToCloudinary(src, file.Filename)
			if err != nil {
				return c.JSON(500, helpers.ResponseFormat(500, "article pic upload error", nil))
			}
		}
		err = ah.srv.UpdateArticle(ToArticleEntity(newArticle), imageURL, dataUser.ID, uint(articleID))
		if err != nil {
			errCode := 500
			if strings.ContainsAny(err.Error(), "tidak ditemukan") {
				errCode = 400
			}
			return c.JSON(errCode, helpers.ResponseFormat(errCode, err.Error(), nil))
		}
		return nil
	}
}
