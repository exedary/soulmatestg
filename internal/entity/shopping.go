package entity

type ShoppingList struct {
	Items      []shoppingItem
	IsFinished bool
}

type shoppingItem struct {
	Description string
	IsChecked   bool
}
