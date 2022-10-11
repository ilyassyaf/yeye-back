package services

import "github.com/ilyassyaf/yeyebackend/models"

type TokenService interface {
	GetAllByCategory() ([]models.TokenCategory, error)
	StoreCategory(*models.TokenCategoryStore) (*models.TokenCategory, error)
	GetAll() ([]models.TokenRes, error)
	StoreToken(uint, *models.TokenStore) (*models.TokenRes, error)
	GetByCategory(string) (*models.TokenCategory, error)
	Get(uint) (*models.TokenRes, error)
}
