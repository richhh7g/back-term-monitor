package mongo_document

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Competitor struct {
	ID       primitive.ObjectID  `bson:"_id,omitempty"`
	FoundAt  primitive.Timestamp `bson:"found_at"`
	Device   string              `bson:"device"`
	BrandID  primitive.ObjectID  `bson:"brand_id"`
	Location string              `bson:"location"`
	Domain   string              `bson:"domain"`
}
