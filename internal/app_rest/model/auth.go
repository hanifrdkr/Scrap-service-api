package model

import "time"

type Auth struct {
	UserID    string    `json:"user_id" bson:"user_id,omitempty"`
	Password  string    `json:"password" bson:"password,omitempty" validate:"required,min=6"`
	Email     string    `json:"email" bson:"email,omitempty" validate:"email,required"`
	CreatedAt time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at,omitempty"`
}

type AccessToken struct {
	Token        string    `json:"token" bson:"token,omitempty"`
	RefreshToken string    `json:"refresh_token" bson:"refresh_token,omitempty"`
	IsRevoke     bool      `json:"is_revoke" bson:"is_revoke"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at" bson:"updated_at,omitempty"`
	ExpiredAt    time.Time `json:"expired_at" bson:"expired_at,omitempty"`
	UserID       string    `json:"user_id" bson:"user_id,omitempty"`
}
