package pair

type CreatePairDto struct {
	UserId    string
	InvitedId string
}

type GetByIdQueryDto struct {
	PairId string `uri:"id" binding:"required,uuid"`
}

type PairDto struct {
	PairId      string
	DateCreated string
}
