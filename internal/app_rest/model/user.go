package model

import "time"

type User struct {
	ID        string    `json:"id" bson:"id,omitempty"`
	Name      string    `json:"name" bson:"name,omitempty"`
	Email     string    `json:"email" bson:"email,omitempty"`
	CreatedAt time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at,omitempty"`
}
