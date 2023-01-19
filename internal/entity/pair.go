package entity

import "time"

type Pair struct {
	Id                    string
	DateCreated           time.Time
	ShoppingList          []ShoppingList
	RelationshipArtifacts []artifact
	Acceptance            PairAcceptance
}

type PairAcceptance struct {
	User       User
	IsAccepted bool
}
