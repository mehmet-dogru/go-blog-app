package dto

type UserLogin struct {
	Email    string `json:"email" validate:"required,email,min=1,max=255"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserSignup struct {
	UserLogin
	FirstName string `json:"first_name,omitempty" validate:"required"`
	LastName  string `json:"last_name,omitempty" validate:"required"`
}

type ProfileInfo struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
