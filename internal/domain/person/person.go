package person

import "go.mongodb.org/mongo-driver/bson/primitive"

type Person struct {
	Id         primitive.ObjectID `bson:"_id"`
	Name       string             `bson:"name"`
	Email      string             `bson:"email"`
	ExternalId string             `bson:"external_id"`
}

func NewPerson(name, email, externalId string) *Person {
	return &Person{
		Id:         primitive.NewObjectID(),
		Name:       name,
		Email:      email,
		ExternalId: externalId,
	}
}

func (person *Person) ChangeName(name string) {
	person.Name = name
}
