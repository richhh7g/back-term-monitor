package mongo_document

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Competitor struct {
	ID        primitive.ObjectID  `bson:"_id,omitempty"`
	Term      string              `bson:"term"`
	Device    string              `bson:"device"`
	Domain    string              `bson:"domain"`
	FoundAt   primitive.Timestamp `bson:"found_at"`
	BrandID   primitive.ObjectID  `bson:"brand_id"`
	Location  string              `bson:"location"`
	CreatedAt primitive.Timestamp `bson:"created_at"`
}
