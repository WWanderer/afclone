package form1

import (
	"github.com/WWanderer/afclone/models/common"
)

type Form1 struct {
	ID                        string                  `json:"ID,omitempty" bson:"ID" form:"id"` // TODO switch back to uuid.UUID, problems with encoding/decoding in mongodriver
	Context                   common.Control[Context] `json:"context,omitempty" bson:"context" form:"context"`
	ParticipatingOrganisations common.Control[ParticipatingOrganisations]  `json:"participating_organisation,omitempty" bson:"participating_organisation,omitempty" form:"participating_organisation"`
	Activities                common.Control[Activities] `json:"activities,omitempty" bson:"activities,omitempty" form:"activities"`
	Annexes                   common.Control[Annexes] `json:"annexes,omitempty" bson:"annexes,omitempty" form:"annexes"`
}

type Form1DTO struct {
	ID      string     `json:"_id,omitempty" bson:"_id,omitempty" form:"id"`
	Context ContextDTO `json:"context,omitempty" bson:"context,omitempty" form:"context"`
}

func NewForm1(ID string) *Form1 {
	f1 := new(Form1)

	f1.ID = ID

	return f1
}
