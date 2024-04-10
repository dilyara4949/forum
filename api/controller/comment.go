package controller

import (
	"forum/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	CommentUsecase domain.CommentUsecase
}

func (uc *CommentController) Create(c *gin.Context) {
	postID, err := strToInt(c.Param("post-id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Post id is not correct: " + err.Error()})
		return
	}

	var request domain.Comment
	err = c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Create comment error: " + err.Error()})
		return
	}
	request.PostID = postID
	request.Email = c.GetString("x-email")
	_, err = uc.CommentUsecase.Create(c, &request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Create comment error: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, request)

}

func (uc *CommentController) Get(c *gin.Context) {
	postID, err := strToInt(c.Param("post-id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Post id is not correct: " + err.Error()})
		return
	}

	comments, err := uc.CommentUsecase.Get(c, postID) 
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Get comments error: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, comments)
}

func (uc *CommentController) Delete(c *gin.Context) {
	postID, err := strToInt(c.Param("post-id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Post id is not correct: " + err.Error()})
		return
	}
	email := c.GetString("x-email")

	err = uc.CommentUsecase.Delete(c, postID, email) 
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Delete comment error: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: ("comment is deleted")})
}