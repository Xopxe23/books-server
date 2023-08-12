package domain

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password" db:"password_hash"`
}

type SignUpInput struct {
	Name     string `json:"name" binding:"required,gte2"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte6"`
}

type SignInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
