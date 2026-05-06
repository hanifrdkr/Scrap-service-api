package service_auth

type LoginPayload struct {
	Email    string `json:"email" bson:"email,omitempty" validate:"email"`
	Password string `json:"password" bson:"password,omitempty" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
