// go get github.com/go-playground/validator/v10
package libraries

import (
	"database/sql"

	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/pjw1702/go-auth/config"

	en_translations "github.com/go-playground/validator/v10/translations/en"
)

// import (
// 	"reflect"

// 	"github.com/go-playground/locales/en"
// 	ut "github.com/go-playground/universal-translator"

// 	//"github.com/go-playground/validator"
// 	"github.com/go-playground/validator/v10"
// 	en_translations "github.com/go-playground/validator/v10/translations/en"
// )

// type Validation struct {
// 	validate *validator.Validate
// 	trans    ut.Translator
// }

// func NewValidation() *Validation {
// 	translator := en.New()
// 	uni := ut.New(translator, translator)

// 	trans, _ := uni.GetTranslator("en")

// 	validate := validator.New()
// 	en_translations.RegisterDefaultTranslations(validate, trans)

// 	// register tag label
// 	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
// 		name := field.Tag.Get("label")
// 		return name
// 	})

// 	// creating custom error
// 	// customize the error message
// 	validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
// 		return ut.Add("required", "{0} must be filled", true)
// 	}, func(ut ut.Translator, fe validator.FieldError) string {
// 		t, _ := ut.T("required", fe.Field())
// 		return t
// 	})

// 	return &Validation{
// 		validate: validate,
// 		trans:    trans,
// 	}
// }

// func (v *Validation) Struct(s interface{}) interface{} {
// 	errors := make(map[string]string)

// 	err := v.validate.Struct(s)
// 	if err != nil {
// 		for _, e := range err.(validator.ValidationErrors) {
// 			errors[e.StructField()] = e.Translate(v.trans)
// 		}
// 	}

// 	if len(errors) > 0 {
// 		return errors
// 	}

// 	return nil
// }

type Validation struct {
	conn *sql.DB
}

// Connect to DB for check validation
func NewValidation() *Validation {
	conn, err := config.DBConn()

	if err != nil {
		panic(err)
	}

	return &Validation{
		conn: conn,
	}
}

// Validator initate
func (v *Validation) Init() (*validator.Validate, ut.Translator) {
	// calling package translator
	translator := en.New()
	uni := ut.New(translator, translator)

	trans, _ := uni.GetTranslator("en")

	// register translation (en)
	validate := validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)

	// change the default label
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		labelname := field.Tag.Get("label")
		return labelname
	})

	// each of field message of validatation check: ${your name of label} cannot be empty
	validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} cannot be empty", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	// check validatioin compare DB and check submitted information is unique
	validate.RegisterValidation("isunique", func(fl validator.FieldLevel) bool {
		params := fl.Param()
		split_params := strings.Split(params, "-")

		tableName := split_params[0]
		fieldName := split_params[1]
		fieldValue := fl.Field().String()

		return v.checkIsUnique(tableName, fieldName, fieldValue)
	})

	// each of field message of validatation check: ${your name of label} already in use
	validate.RegisterTranslation("isunique", trans, func(ut ut.Translator) error {
		return ut.Add("isunique", "{0} already in use", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("isunique", fe.Field())
		return t
	})

	return validate, trans
}

func (v *Validation) Struct(s interface{}) interface{} {
	validate, trans := v.Init()

	vErrors := make(map[string]interface{})

	err := validate.Struct(s)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			vErrors[e.StructField()] = e.Translate(trans)
		}
	}

	if len(vErrors) > 0 {
		return vErrors
	}

	return nil
}

// Function to check unique
func (v *Validation) checkIsUnique(tableName, fieldName, fieldValue string) bool {

	row, _ := v.conn.Query("select "+fieldName+" from "+tableName+" where "+fieldName+" = ?", fieldValue)
	// select email from users where email = "pjw1702@huevertech.com"

	defer row.Close()

	var result string
	for row.Next() {
		row.Scan(&result)
	}

	// email@huevertech.com

	// It must be an email that does not exist in the users DB
	return result != fieldValue
}
