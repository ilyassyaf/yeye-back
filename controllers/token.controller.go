package controllers

import (
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ilyassyaf/yeyebackend/config"
	"github.com/ilyassyaf/yeyebackend/models"
	"github.com/ilyassyaf/yeyebackend/services"
	"go.mongodb.org/mongo-driver/mongo"
)

type TokenCotroller struct {
	tokenService   services.TokenService
	counterService services.CounterService
}

func NewTokenController(tokenService services.TokenService, counterService services.CounterService) TokenCotroller {
	return TokenCotroller{tokenService, counterService}
}

func (tc *TokenCotroller) StoreCategory(c *gin.Context) {
	var cat *models.TokenCategoryStore

	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	newCat, err := tc.tokenService.StoreCategory(cat)
	if err != nil {
		if strings.Contains(err.Error(), "Category already exist") {
			c.JSON(http.StatusConflict, gin.H{"status": "error", "message": err.Error()})
			return
		}

		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": newCat})
}

func (tc *TokenCotroller) GetAllByCategory(c *gin.Context) {
	tokenList, err := tc.tokenService.GetAllByCategory()
	if err != nil && err != mongo.ErrNoDocuments {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"token_list": tokenList}})
}

func (tc *TokenCotroller) GetAll(c *gin.Context) {
	tokenList, err := tc.tokenService.GetAll()
	if err != nil && err != mongo.ErrNoDocuments {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"token_list": tokenList}})
}

func (tc *TokenCotroller) GetByCategory(c *gin.Context) {
	category := c.Query("category")
	if category == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Please provide param:'category'"})
		return
	}
	result, err := tc.tokenService.GetByCategory(category)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": result})
}

func (tc *TokenCotroller) Store(c *gin.Context) {
	id, err := tc.counterService.GetNextSequence("token")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var token *models.TokenStore

	if err := c.ShouldBind(&token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	newToken, err := tc.tokenService.StoreToken(id.SequenceValue, token)
	if err != nil {
		if strings.Contains(err.Error(), "Token already exist") {
			c.JSON(http.StatusConflict, gin.H{"status": "error", "message": err.Error()})
			return
		}

		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": newToken})
}

func (tc *TokenCotroller) StoreTokenImage(c *gin.Context) {
	config, err := config.LoadConfig("../")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Could not load config",
		})
		return
	}

	file, err := c.FormFile("image_file")

	// The file cannot be received.
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "No file is received",
		})
		return
	}

	// Retrieve file information
	extension := filepath.Ext(file.Filename)
	// Generate random file name for the new uploaded file so it doesn't override the old file with same name
	newFileName := uuid.New().String() + extension

	// The file is received, so let's save it
	if err := c.SaveUploadedFile(file, "assets/"+newFileName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "fail",
			"message": "Unable to save the file: " + err.Error(),
		})
		return
	}

	// File saved successfully. Return proper result
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"path":   config.BaseUrl + "assets/" + newFileName,
	})
}

func (tc *TokenCotroller) Get(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Query("id"), 10, 32)
	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Please provide param:'id'"})
		return
	}
	tokenRes, err := tc.tokenService.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": tokenRes})
}

func (tc *TokenCotroller) GetMetadata(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Please provide param:'id'"})
		return
	}
	metadataRes, err := tc.tokenService.GetMetadata(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
		return
	}

	c.JSON(http.StatusOK, metadataRes)
}
