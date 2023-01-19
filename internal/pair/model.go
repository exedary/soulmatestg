package pair

type createPairDto struct {
	UserId    string `json:"userId"`
	InvitedId string `json:"invitedId"`
}

type pairDto struct {
	PairId      string `json:"pairId"`
	DateCreated string `json:"dateCreated"`
}
