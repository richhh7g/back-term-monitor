package mongo_document

import "go.mongodb.org/mongo-driver/bson/primitive"

type Alarm struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Email string             `bson:"email,omitempty"`
	Tags  []string           `bson:"tags,omitempty"`
}
