// go get golang.org/x/crypto/bcrypt
package controllers

import (
	"errors"
	"html/template"
	"net/http"

	"github.com/pjw1702/go-auth/config"
	entites "github.com/pjw1702/go-auth/entities"
	"github.com/pjw1702/go-auth/models"
	"golang.org/x/crypto/bcrypt"
)

type UserInput struct {
	Username string
	Password string
}

var userModel = models.NewUserModel()

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

func Logout(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Store.Get(r, config.SESSOION_ID)

	// delete session
	session.Options.MaxAge = -1
	session.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
