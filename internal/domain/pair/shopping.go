package pair

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ShoppingList struct {
	Id          primitive.ObjectID `bson:"_id"`
	Items       []ShoppingItem
	DateCreated time.Time
	IsFinished  bool
}

type ShoppingItem struct {
	Id          primitive.ObjectID `bson:"_id"`
	Description string             `bson:"description"`
	IsChecked   bool               `bson:"isChecked"`
}

func newShoppingList(items []ShoppingItem) *ShoppingList {
	return &ShoppingList{
		Id:          primitive.NewObjectID(),
		DateCreated: time.Now().UTC(),
		IsFinished:  false,
		Items:       items,
	}
}

func (shoppingList *ShoppingList) Append(item *ShoppingItem) {
	shoppingList.Items = append(shoppingList.Items, *item)
}

func (shoppingList *ShoppingList) MarkFinished() {
	shoppingList.IsFinished = true
}
