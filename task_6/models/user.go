package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username string             `bson:"username" json:"username" binding:"required"`
	Password string             `bson:"password" json:"password,omitempty"`
	Role     string             `bson:"role" json:"role,omitempty"` // "admin" or "user"
}
