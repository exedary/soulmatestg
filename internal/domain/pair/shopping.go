package pair

import (
	"time"

	"github.com/exedary/soulmates/internal/uuid"
)

type ShoppingList struct {
	Id          string
	Items       []ShoppingItem
	DateCreated time.Time
	IsFinished  bool
}

type ShoppingItem struct {
	Id          string
	Description string
	IsChecked   bool
}

func newShoppingList(items []ShoppingItem) *ShoppingList {
	return &ShoppingList{
		Id:          uuid.Generate(),
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
