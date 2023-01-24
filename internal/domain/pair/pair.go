package pair

import (
	"time"

	"github.com/exedary/soulmates/internal/domain/person"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pair struct {
	Id           primitive.ObjectID   `bson:"_id"`
	DateCreated  time.Time            `bson:"dateCreated"`
	ShoppingList []ShoppingList       `bson:"ShoppingList"`
	Participants []primitive.ObjectID `bson:"Participants"`
}

func New(createdBy *person.Person) *Pair {
	participants := make([]primitive.ObjectID, 2)
	participants[0] = createdBy.Id

	return &Pair{
		Id:           primitive.NewObjectIDFromTimestamp(time.Now()),
		DateCreated:  time.Now().UTC(),
		Participants: participants,
	}
}

func (pair *Pair) AcceptInvitation(invited *person.Person) {
	pair.Participants[1] = invited.Id
}

func (pair *Pair) AttachShoppingList(items []ShoppingItem) {
	pair.ShoppingList = append(pair.ShoppingList, *newShoppingList(items))
}
