package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type TokenCategory struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Category string             `json:"category" bson:"category"`
	Token    []Token            `json:"token,omitempty" bson:"token,omitempty"`
}

type Token struct {
	ID       uint          `json:"_id" bson:"_id"`
	Image    string        `json:"image_name" bson:"image_name"`
	Metadata TokenMetadata `json:"metadata" bson:"metadata"`
}

type TokenMetadata struct {
	Name            string           `json:"name" bson:"name"`
	Description     string           `json:"description" bson:"description"`
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
	Value       string `json:"value" bson:"value"`
}
