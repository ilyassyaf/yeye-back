package repository

import (
	"context"

	"github.com/ilyassyaf/yeyebackend/models"
	"github.com/ilyassyaf/yeyebackend/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CounterServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewCounterServiceImpl(collection *mongo.Collection, ctx context.Context) services.CounterService {
	return &CounterServiceImpl{collection, ctx}
}

func (cs *CounterServiceImpl) GetNextSequence(id string) (*models.GetCounter, error) {
	var sequence *models.GetCounter

	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$inc", Value: bson.D{{Key: "sequence_value", Value: 1}}}}
	opt := options.FindOneAndUpdate().SetUpsert(true)

	err := cs.collection.FindOneAndUpdate(cs.ctx, filter, update, opt).Decode(&sequence)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &models.GetCounter{}, err
		}
		return nil, err
	}

	return sequence, nil
}
