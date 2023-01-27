package entities

type Patient struct {
	// label: reflect name of label when get message of validate error instead of name of variable
	Id             int64
	Full_Name      string `validate:"required" label:"Full Name"`
	Number_Patient string `validate:"required" label:"Patient Number"`
	Gender         string `validate:"required"`
	Birthplace     string `validate:"required"`
	Birthday       string `validate:"required"`
	Address        string `validate:"required"`
	Number_HP      string `validate:"required" label:"Phone Number"`
}
