package controller

import (
	"fmt"
	"forum/internal/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	PostUsecase domain.PostUsecase
}

func (uc *PostController) Create(c *gin.Context) {
	var request domain.PostRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Create post error: " + err.Error()})
		return
	}
	request.Email = c.GetString("x-email")
	_, err = uc.PostUsecase.Create(c, &request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Create post error: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, request)

}

func (uc *PostController) GetOwn(c *gin.Context) {
	email := c.GetString("x-email")

	post, err := uc.PostUsecase.GetOwn(c, email)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Get posts error: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}

func (uc *PostController) GetAll(c *gin.Context) {

	posts, err := uc.PostUsecase.GetAll(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Get all posts error: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)
}

func (uc *PostController) Delete(c *gin.Context) {
	id, err := strToInt(c.Param("post-id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Post id is not correct: " + err.Error()})
		return
	}
	err = uc.PostUsecase.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Delete post error: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: fmt.Sprintf("post with id: %d is deleted", id)})
}

func strToInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	return  i, err
}
