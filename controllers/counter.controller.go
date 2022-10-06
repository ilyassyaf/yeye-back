package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilyassyaf/yeyebackend/services"
	"go.mongodb.org/mongo-driver/mongo"
)

type CounterController struct {
	counterService services.CounterService
}

func NewCounterController(counterService services.CounterService) CounterController {
	return CounterController{counterService}
}

func (cc *CounterController) GetNextSequence(c *gin.Context) {
	seqName := c.Query("seq_name")
	if seqName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "param: seq_name is required"})
		return
	}
	sequence, err := cc.counterService.GetNextSequence(seqName)
	if err != nil && err != mongo.ErrNoDocuments {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"sequence": sequence}})
}
