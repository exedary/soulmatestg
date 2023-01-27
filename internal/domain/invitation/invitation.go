package invitation

import (
	"github.com/exedary/soulmates/internal/domain/person"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invitation struct {
	Id         primitive.ObjectID
	From       primitive.ObjectID
	AcceptedBy primitive.ObjectID
	Accepted   bool
}

func New(from *person.Person) *Invitation {
	return &Invitation{
		Id:       primitive.NewObjectID(),
		From:     from.Id,
		Accepted: false,
	}
}

func (invitation *Invitation) Accept(acceptedBy *person.Person) {
	invitation.AcceptedBy = acceptedBy.Id
	invitation.Accepted = true
}
