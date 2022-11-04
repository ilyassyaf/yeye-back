package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type TokenCategoryStore struct {
	Category string `json:"category" bson:"category" binding:"required"`
}

type TokenCategory struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Category string             `json:"category" bson:"category"`
	Token    []Token            `json:"token,omitempty" bson:"token,omitempty"`
}

type TokenStore struct {
	ID       uint               `json:"_id" bson:"_id"`
	Category primitive.ObjectID `json:"category,omitempty" bson:"category,omitempty"`
	Metadata TokenMetadataStore `json:"metadata" bson:"metadata" binding:"required"`
}

type Token struct {
	ID       uint               `json:"_id" bson:"_id"`
	Image    string             `json:"image_name" bson:"image_name"`
	Category primitive.ObjectID `json:"category,omitempty" bson:"category,omitempty"`
	Metadata TokenMetadata      `json:"metadata,omitempty" bson:"metadata,omitempty"`
}

type TokenRes struct {
	ID       uint          `json:"_id" bson:"_id"`
	Image    string        `json:"image_name" bson:"image_name"`
	Category string        `json:"category,omitempty" bson:"category,omitempty"`
	Metadata TokenMetadata `json:"metadata,omitempty" bson:"metadata,omitempty"`
}

type TokenMetadataStore struct {
	Name            string           `json:"name" bson:"name" binding:"required"`
	Description     string           `json:"description" bson:"description" binding:"required"`
	Image           string           `json:"image" bson:"image" binding:"required"`
	ExternalUrl     string           `json:"external_url" bson:"external_url"`
	YoutubeUrl      string           `json:"youtube_url" bson:"youtube_url"`
	BackgroundColor string           `json:"background_color" bson:"background_color"`
	AnimationUrl    string           `json:"animation_url" bson:"animation_url"`
	Attributes      []TokenAttribute `json:"attributes" bson:"attributes"`
}

type TokenMetadata struct {
	Name            string           `json:"name,omitempty" bson:"name,omitempty"`
	Description     string           `json:"description,omitempty" bson:"description,omitempty"`
	Image           string           `json:"image,omitempty" bson:"image,omitempty"`
	ExternalUrl     string           `json:"external_url,omitempty" bson:"external_url,omitempty"`
	YoutubeUrl      string           `json:"youtube_url,omitempty" bson:"youtube_url,omitempty"`
	BackgroundColor string           `json:"background_color,omitempty" bson:"background_color,omitempty"`
	AnimationUrl    string           `json:"animation_url,omitempty" bson:"animation_url,omitempty"`
	Attributes      []TokenAttribute `json:"attributes,omitempty" bson:"attributes,omitempty"`
}

type TokenAttribute struct {
	DisplayType string `json:"display_type,omitempty" bson:"display_type,omitempty"`
	TraitType   string `json:"trait_type,omitempty" bson:"trait_type,omitempty"`
	Value       string `json:"value,omitempty" bson:"value,omitempty"`
}
