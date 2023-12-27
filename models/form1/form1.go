package form1

import (
	"github.com/WWanderer/afclone/models/common"
)

type Form1 struct {
	ID      string                  `json:"ID,omitempty" bson:"ID" form:"id"` // TODO switch back to uuid.UUID, problems with encoding/decoding in mongodriver
	Context common.Control[Context] `json:"context,omitempty" bson:"context" form:"context"`
}

type Form1DTO struct {
	ID      string     `json:"_id,omitempty" bson:"_id,omitempty" form:"id"`
	Context ContextDTO `json:"context,omitempty" bson:"context,omitempty" form:"context"`
}

func NewForm1(ID string) *Form1 {
	f1 := new(Form1)

	f1.ID = ID
	f1.Context.Section = true

	return f1
}