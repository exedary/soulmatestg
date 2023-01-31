package invitation

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func byIdSpec(id primitive.ObjectID) primitive.M {
	filter := bson.M{
		"_id": id,
	}

	return filter
}
