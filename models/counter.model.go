package models

type GetCounter struct {
	ID            string `json:"id" bson:"_id" binding:"required"`
	SequenceValue uint   `json:"sequence_value" bson:"sequence_value"`
}
