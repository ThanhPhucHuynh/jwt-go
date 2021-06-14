package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type UserGet struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"` // tag golang 
	Email       string             `json:"email" bson:"email"`
	DisplayName string             `json:"displayName" bson:"displayName"`
}
