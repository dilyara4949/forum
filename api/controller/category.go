package controller

import (
	"fmt"
	"forum/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	CategoryUsecase domain.CategoryUsecase
}

func (uc *CategoryController) Create(c *gin.Context) {
	var request domain.Category
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Create category error: " + err.Error()})
		return
	}

	_, err = uc.CategoryUsecase.Create(c, &request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Create category error: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, request)

}

func (uc *CategoryController) Get(c *gin.Context) {
	name := c.Param("name")

	category, err := uc.CategoryUsecase.Get(c, name) 
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Get category error: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, category)
}

func (uc *CategoryController) GetAll(c *gin.Context) {

	categories, err := uc.CategoryUsecase.GetAll(c) 
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Get all category error: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}

func (uc *CategoryController) Delete(c *gin.Context) {
	name := c.Param("name")

	err := uc.CategoryUsecase.Delete(c, name) 
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Delete category error: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: fmt.Sprintf("category: %s deleted", name)})
}