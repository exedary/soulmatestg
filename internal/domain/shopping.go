package domain

import "time"

type ShoppingList struct {
	Items       []shoppingItem
	IsFinished  bool
	DateCreated time.Time
}

type shoppingItem struct {
	Description string
	Checked     bool
}
