package pair

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

func byPersonIdSpec(personId primitive.ObjectID) primitive.M {
	filter := bson.M{
		"participants": personId,
	}

	return filter
}
