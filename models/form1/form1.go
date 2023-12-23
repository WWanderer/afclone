package form1

import (
	"github.com/WWanderer/afclone/models/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Form1 struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"id,omitempty" form:"id"`
	Context common.Control[Context] `_json:"context,omitempty" bson:"context,omitempty" form:"context"`
}

type Form1DTO struct {
	ID      string `json:"_id,omitempty" bson:"_id,omitempty" form:"id"`
	Context ContextDTO         `json:"context,omitempty" bson:"context,omitempty" form:"context"`
}
