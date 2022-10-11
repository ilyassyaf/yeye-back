package repository

import (
	"context"
	"errors"

	"github.com/ilyassyaf/yeyebackend/models"
	"github.com/ilyassyaf/yeyebackend/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TokenServiceImpl struct {
	db  *mongo.Database
	ctx context.Context
}

func NewTokenServiceImpl(db *mongo.Database, ctx context.Context) services.TokenService {
	return &TokenServiceImpl{db, ctx}
}

func (ts *TokenServiceImpl) StoreCategory(newCategory *models.TokenCategoryStore) (*models.TokenCategory, error) {
	res, err := ts.db.Collection("token_categories").InsertOne(ts.ctx, &newCategory)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("Category already exist")
		}
		return nil, err
	}

	// Create a unique index for the email field
	opt := options.Index()
	opt.SetUnique(true)
	index := mongo.IndexModel{Keys: bson.M{"category": 1}, Options: opt}

	if _, err := ts.db.Collection("token_categories").Indexes().CreateOne(ts.ctx, index); err != nil {
		return nil, errors.New("Could not create index for category")
	}

	var newCat *models.TokenCategory
	query := bson.M{"_id": res.InsertedID}

	err = ts.db.Collection("token_categories").FindOne(ts.ctx, query).Decode(&newCat)
	if err != nil {
		return nil, err
	}

	return newCat, nil
}

func (ts *TokenServiceImpl) GetAllByCategory() ([]models.TokenCategory, error) {
	var tokenList []models.TokenCategory
	qry := bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: "token"},
		{Key: "localField", Value: "_id"},
		{Key: "foreignField", Value: "category"},
		{Key: "as", Value: "token"},
	}}}

	curr, err := ts.db.Collection("token_categories").Aggregate(ts.ctx, mongo.Pipeline{qry})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return []models.TokenCategory{}, err
		}
		return nil, err
	}

	if err := curr.All(ts.ctx, &tokenList); err != nil {
		return nil, err
	}

	return tokenList, nil
}

func (ts *TokenServiceImpl) GetAll() ([]models.TokenRes, error) {
	var tokenList []models.TokenRes
	lookup := bson.D{
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "token_categories"},
			{Key: "localField", Value: "category"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "category"},
		}},
	}
	unwind := bson.D{
		{Key: "$unwind", Value: bson.D{
			{Key: "path", Value: "$category"},
		}},
	}
	cat := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "category", Value: "$category.category"},
		}},
	}

	curr, err := ts.db.Collection("token").Aggregate(ts.ctx, mongo.Pipeline{lookup, unwind, cat})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return []models.TokenRes{}, err
		}
		return nil, err
	}

	if err := curr.All(ts.ctx, &tokenList); err != nil {
		return nil, err
	}

	return tokenList, nil
}

func (ts *TokenServiceImpl) StoreToken(nextID uint, newToken *models.TokenStore) (*models.TokenRes, error) {
	newToken.ID = nextID
	res, err := ts.db.Collection("token").InsertOne(ts.ctx, &newToken)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("Token already exist")
		}
		return nil, err
	}

	var insertToken *models.TokenRes
	query := bson.M{"_id": res.InsertedID}

	err = ts.db.Collection("token").FindOne(ts.ctx, query).Decode(&insertToken)
	if err != nil {
		return nil, err
	}

	return insertToken, nil
}

func (ts *TokenServiceImpl) GetByCategory(cat string) (*models.TokenCategory, error) {
	var tokenList *models.TokenCategory

	match := bson.D{
		{Key: "$match", Value: bson.D{
			{Key: "category", Value: cat},
		}},
	}
	lookup := bson.D{
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "token"},
			{Key: "localField", Value: "_id"},
			{Key: "foreignField", Value: "category"},
			{Key: "as", Value: "token"},
		}},
	}
	curr, err := ts.db.Collection("token_categories").Aggregate(ts.ctx, mongo.Pipeline{match, lookup})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &models.TokenCategory{}, err
		}
		return nil, err
	}

	curr.Next(ts.ctx)
	if err := curr.Decode(&tokenList); err != nil {
		return nil, err
	}

	return tokenList, nil
}

func (ts *TokenServiceImpl) Get(id uint) (*models.TokenRes, error) {
	var tokenRes *models.TokenRes
	match := bson.D{
		{Key: "$match", Value: bson.D{
			{Key: "_id", Value: id},
		}},
	}
	lookup := bson.D{
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "token_categories"},
			{Key: "localField", Value: "category"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "category"},
		}},
	}
	unwind := bson.D{
		{Key: "$unwind", Value: bson.D{
			{Key: "path", Value: "$category"},
		}},
	}
	cat := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "category", Value: "$category.category"},
		}},
	}

	curr, err := ts.db.Collection("token").Aggregate(ts.ctx, mongo.Pipeline{match, lookup, unwind, cat})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &models.TokenRes{}, err
		}
		return nil, err
	}

	curr.Next(ts.ctx)
	if err := curr.Decode(&tokenRes); err != nil {
		return nil, err
	}

	return tokenRes, nil
}
