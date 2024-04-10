package controller

import (
	"forum/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LikeController struct {
	LikeUsecase domain.LikeUsecase
}

func (uc *LikeController) Create(c *gin.Context) {
	postID, err := strToInt(c.Param("post-id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Post id is not correct: " + err.Error()})
		return
	}

	var request domain.Like
	
	request.PostID = postID
	email := c.GetString("x-email")
	_, err = uc.LikeUsecase.Create(c, &request, email)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Create like error: " + err.Error()})
		return
	}

	like, err := uc.LikeUsecase.Get(c, postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Create like error: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, like)

}

func (uc *LikeController) Get(c *gin.Context) {
	postID, err := strToInt(c.Param("post-id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Post id is not correct: " + err.Error()})
		return
	}

	likes, err := uc.LikeUsecase.Get(c, postID) 
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Get likes error: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, likes)
}

func (uc *LikeController) Delete(c *gin.Context) {
	postID, err := strToInt(c.Param("post-id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Post id is not correct: " + err.Error()})
		return
	}
	email := c.GetString("x-email")

	likes, err := uc.LikeUsecase.Delete(c, postID, email) 
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Delete like error: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, likes)
}