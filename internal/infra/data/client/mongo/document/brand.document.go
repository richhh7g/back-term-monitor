package mongo_document

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	BrandErrorStatus   = "error"
	BrandSuccessStatus = "success"
	BrandPendingStatus = "pending"
)

type Brand struct {
	ID         primitive.ObjectID   `bson:"_id,omitempty"`
	CreatedAt  primitive.Timestamp  `bson:"created_at"`
	Email      string               `bson:"email"`
	Status     string               `bson:"status"`
	Results    []primitive.ObjectID `bson:"results"`
	BrandTerms []string             `bson:"brand_terms"`
}
