package request

type User struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=3,max=100,containsany=!@#$%*"`
	Name     string `json:"name" binding:"required,min=3,max=50"`
	Age      int8   `json:"age" binding:"required,min=18,max=140"`
}
