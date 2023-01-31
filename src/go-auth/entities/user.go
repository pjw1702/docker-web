package entites

// validate property: github.com/go-playground/validator
type User struct {
	Id               int64
	Full_Name        string `validate:"required" label:"Full Name"`
	Email            string `validate:"required,email,isunique=users-email"`
	Username         string `validate:"required,gte=3,isunique=users-username"`
	Password         string `validate:"required,gte=6"`
	Confirm_Password string `validate:"required,eqfield=Password" label:"Confirm Password"`
}
