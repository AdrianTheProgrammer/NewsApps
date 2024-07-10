package handlers

import (
	"newsapps/internal/features/comments"
	"newsapps/internal/helpers"
	"newsapps/internal/utils"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type CommentsHand struct {
	srv comments.Services
}

func NewCommentsHand(s comments.Services) comments.Handlers {
	return &CommentsHand{
		srv: s,
	}
}

func (ah *CommentsHand) ShowAllComments() echo.HandlerFunc {
	return func(c echo.Context) error {
		articleID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(400, helpers.ResponseFormat(400, "article id Error", nil))
		}
		responseData, err := ah.srv.ShowAllComments(uint(articleID))
		if err != nil {
			errCode := 500
			if strings.ContainsAny(err.Error(), "tidak ditemukan") {
				errCode = 400
			}
			return c.JSON(errCode, helpers.ResponseFormat(errCode, err.Error(), nil))
		}
		return c.JSON(200, helpers.ResponseFormat(200, "data comments berhasil diambil", responseData))
	}
}

func (ah *CommentsHand) CreateComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		var dataUser = utils.NewTokenUtil().DecodeToken(c.Get("user").(*jwt.Token))
		articleID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(400, helpers.ResponseFormat(400, "article id Error", nil))
		}
		var newComment NewCommentsRequest
		err = c.Bind(&newComment)
		if err != nil {
			return c.JSON(500, helpers.ResponseFormat(500, "comment data Error", nil))
		}

		err = ah.srv.CreateComment(ToCommentEntity(newComment), uint(articleID), dataUser.ID)
		if err != nil {
			return c.JSON(500, helpers.ResponseFormat(500, "comment post error", nil))
		}
		return nil
	}

}

func (ah *CommentsHand) DeleteComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		var dataUser = utils.NewTokenUtil().DecodeToken(c.Get("user").(*jwt.Token))
		commentID, err := strconv.Atoi(c.Param("cid"))
		if err != nil {
			return c.JSON(400, helpers.ResponseFormat(400, "comment id Error", nil))
		}

		err = ah.srv.DeleteComment(uint(commentID), dataUser.ID)
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

func (ah *CommentsHand) ReadComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		commentID, err := strconv.Atoi(c.Param("cid"))
		if err != nil {
			return c.JSON(400, helpers.ResponseFormat(400, "comment id Error", nil))
		}

		responseData, err := ah.srv.ReadComment(uint(commentID))
		if err != nil {
			errCode := 500
			if strings.ContainsAny(err.Error(), "tidak ditemukan") {
				errCode = 400
			}
			return c.JSON(errCode, helpers.ResponseFormat(errCode, err.Error(), nil))
		}
		return c.JSON(200, helpers.ResponseFormat(200, "data comment berhasil diambil", responseData))
	}
}

func (ah *CommentsHand) UpdateComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		var dataUser = utils.NewTokenUtil().DecodeToken(c.Get("user").(*jwt.Token))
		commentID, err := strconv.Atoi(c.Param("cid"))
		if err != nil {
			return c.JSON(400, helpers.ResponseFormat(400, "comment id Error", nil))
		}
		var newComment NewCommentsRequest
		err = c.Bind(&newComment)
		if err != nil {
			return c.JSON(400, helpers.ResponseFormat(400, "comment data Error", nil))
		}

		err = ah.srv.UpdateComment(ToCommentEntity(newComment), dataUser.ID, uint(commentID))
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
