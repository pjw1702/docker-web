// go get golang.org/x/crypto/bcrypt
package controllers

import (
	"errors"
	"html/template"
	"net/http"

	//"reflect"
	//"strings"

	//"github.com/go-playground/locales/en"
	//ut "github.com/go-playground/universal-translator"

	"github.com/pjw1702/go-auth/config"
	entites "github.com/pjw1702/go-auth/entities"
	"github.com/pjw1702/go-auth/libraries"
	"github.com/pjw1702/go-auth/models"
	"golang.org/x/crypto/bcrypt"
	//en_translations "github.com/go-playground/validator/v10/translations/en"
)

type UserInput struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}

var userModel = models.NewUserModel()
var validation = libraries.NewValidation()

func Index(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Store.Get(r, config.SESSOION_ID)

	if len(session.Values) == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {

			data := map[string]interface{}{
				"full_name": session.Values["full_name"],
			}

			temp, _ := template.ParseFiles("views/index.html")
			temp.Execute(w, data)
		}
	}

}

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		temp, _ := template.ParseFiles("views/login.html")
		temp.Execute(w, nil)

	} else if r.Method == http.MethodPost {
		// login process
		r.ParseForm()
		UserInput := &UserInput{
			Username: r.Form.Get("username"),
			Password: r.Form.Get("password"),
		}

		// Check validation when login

		errorMessages := validation.Struct(UserInput)

		// perform a validation process
		if errorMessages != nil {

			//fmt.Println(errorMessages)
			data := map[string]interface{}{
				"validation": errorMessages,
			}

			temp, _ := template.ParseFiles("views/login.html")
			temp.Execute(w, data)
		} else {

			var user entites.User
			userModel.Where(&user, "username", UserInput.Username)

			var message error
			if user.Username == "" {
				// not found in database
				message = errors.New("Wrong Username or Password!")
			} else {
				// password checking
				errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(UserInput.Password))
				if errPassword != nil {
					message = errors.New("Wrong Username or Password!")
				}

			}

			if message != nil {

				data := map[string]interface{}{
					"error": message,
				}

				temp, _ := template.ParseFiles("views/login.html")
				temp.Execute(w, data)
			} else {
				// set session
				session, _ := config.Store.Get(r, config.SESSOION_ID)

				session.Values["loggedIn"] = true
				session.Values["email"] = user.Email
				session.Values["username"] = user.Username
				session.Values["full_name"] = user.Full_Name

				session.Save(r, w)

				http.Redirect(w, r, "/", http.StatusSeeOther)
			}

		}
	}

}

func Logout(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Store.Get(r, config.SESSOION_ID)

	// delete session
	session.Options.MaxAge = -1
	session.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func Register(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		temp, _ := template.ParseFiles("views/register.html")
		temp.Execute(w, nil)

	} else if r.Method == http.MethodPost {
		// perform a registration process

		// take a form instruction
		r.ParseForm()

		user := entites.User{
			Full_Name:        r.Form.Get("full_name"),
			Email:            r.Form.Get("email"),
			Username:         r.Form.Get("username"),
			Password:         r.Form.Get("password"),
			Confirm_Password: r.Form.Get("confirm_password"),
		}

		// errorMessages := make(map[string]interface{})

		// if user.Full_Name == "" {
		// 	errorMessages["Full_Name"] = "Full Name must be filled in"
		// }
		// if user.Email == "" {
		// 	errorMessages["Email"] = "Email must be filled in"
		// }
		// if user.Username == "" {
		// 	errorMessages["Username"] = "Username must be filled in"
		// }
		// if user.Password == "" {
		// 	errorMessages["Password"] = "Password must be filled in"
		// }
		// if user.Confirm_Password == "" {
		// 	errorMessages["Confirm_Password"] = "Confrim Password must be filled in"
		// } else {
		// 	if user.Confirm_Password != user.Password {
		// 		errorMessages["Confirm_Password"] = "Password configuration mismatch"
		// 	}
		// }

		// if len(errorMessages) > 0 {
		// 	// form validation failed

		// 	data := map[string]interface{}{
		// 		"validation": errorMessages,
		// 	}

		// 	temp, _ := template.ParseFiles("views/register.html")
		// 	temp.Execute(w, data)
		// } else {
		// 	// convert inserted passord to hash password using bcrypt
		// 	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		// 	user.Password = string(hashPassword)

		// 	// insert to DB
		// 	_, err := userModel.Create(user)

		// 	var message string
		// 	if err != nil {
		// 		message = "perform a registration process: " + message
		// 	} else {
		// 		message = "Registration succeeded, please login"
		// 	}

		// 	data := map[string]interface{}{
		// 		"message": message,
		// 	}

		// 	temp, _ := template.ParseFiles("views/register.html")
		// 	temp.Execute(w, data)
		// }

		// Check validate

		// calling package translator
		// translator := en.New()
		// uni := ut.New(translator, translator)

		// trans, _ := uni.GetTranslator("en")

		// // register translation (en)
		// validate := validator.New()
		// en_translations.RegisterDefaultTranslations(validate, trans)

		// // change the default label
		// validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		// 	labelname := field.Tag.Get("label")
		// 	return labelname
		// })

		// // each of field message of validatation check: ${your name of label} cannot be empty
		// validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		// 	return ut.Add("required", "{0} cannot be empty", true)
		// }, func(ut ut.Translator, fe validator.FieldError) string {
		// 	t, _ := ut.T("required", fe.Field())
		// 	return t
		// })

		// validate.RegisterValidation("isunique", func(fl validator.FieldLevel) bool {
		// 	params := fl.Param()
		// 	split_params := strings.Split(params, "-")

		// 	tableName := split_params[0]
		// 	fieldName := split_params[1]
		// 	fieldValue := fl.Field().String()

		// 	return checkIsUnique(tableName, fieldName, fieldValue)
		// })

		// // each of field message of validatation check: ${your name of label} already in use
		// validate.RegisterTranslation("isunique", trans, func(ut ut.Translator) error {
		// 	return ut.Add("isunique", "{0} already in use", true)
		// }, func(ut ut.Translator, fe validator.FieldError) string {
		// 	t, _ := ut.T("isunique", fe.Field())
		// 	return t
		// })

		errorMessages := validation.Struct(user)

		// perform a validation process
		if errorMessages != nil {

			//fmt.Println(errorMessages)
			data := map[string]interface{}{
				"validation": errorMessages,
				"user":       user,
			}

			temp, _ := template.ParseFiles("views/register.html")
			temp.Execute(w, data)
		} else {

			// hashPassword
			hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			user.Password = string(hashPassword)

			// insert into DB
			userModel.Create(user)

			data := map[string]interface{}{
				"message": "Registration succeeded",
			}
			temp, _ := template.ParseFiles("views/register.html")
			temp.Execute(w, data)

		}
	}

}
