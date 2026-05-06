package service_auth

type RegisterPayload struct {
	Name     string `json:"name" bson:"name,omitempty"`
	Email    string `json:"email" bson:"email,omitempty"`
	Password string `json:"password" bson:"password,omitempty"`
}
