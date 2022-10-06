package services

import "github.com/ilyassyaf/yeyebackend/models"

type CounterService interface {
	GetNextSequence(string) (*models.GetCounter, error)
}
