package authcontroller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pjw1702/go-jwt-mux/config"
	"github.com/pjw1702/go-jwt-mux/helper"
	"github.com/pjw1702/go-jwt-mux/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// take json input
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		//response := map[string]string{"message": err.Error()}
		response := make(map[string]string)
		response["message"] = err.Error()
		helper.ResponseJSON(w, http.StatusBadRequest, response)
	}
	defer r.Body.Close()

	// get user data by username
	var user models.User
	if err := models.DB.Where("username = ?", userInput.Username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "Wrong username or password"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{"message": err.Error()}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
	}

	// check password valid
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	// jwt token generation process
	// ExpiresAT: jwt.NewNumericDate(expTime),
	expTime := time.Now().Add(time.Minute * 1)
	claims := &config.JWTClaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "go-jwt-mux",
			//Subject:   "",
			//Audience:  []string{},
			ExpiresAt: &jwt.NumericDate{Time: expTime},
			//NotBefore: &jwt.NumericDate{},
			//IssuedAt:  &jwt.NumericDate{},
			//ID:        "",
		},
	}

	// declare algorithm to be used for signing
	// if set jwt.SigningMethodES256, key will be not valid at try to login
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sigined token
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	// set token to cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	response := map[string]string{"message": "login successful"}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func Register(w http.ResponseWriter, r *http.Request) {

	// take json input
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		//response := map[string]string{"message": err.Error()}
		response := make(map[string]string)
		response["message"] = err.Error()
		helper.ResponseJSON(w, http.StatusBadRequest, response)
	}
	defer r.Body.Close()

	// hash pass using bcrypt
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	// insert key database
	if err := models.DB.Create(&userInput).Error; err != nil {

		//response := map[string]string{"message": err.Error()}

		response := make(map[string]string)
		response["message"] = err.Error()
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
	}

	//response, _ := json.Marshal(map[string]string{"message": "success"})
	//w.Header().Add("Content-Type", "application/json")
	//w.WriteHeader(http.StatusOK)
	//w.Write(response)

	//response := map[string]string{"message": "success"}

	response := make(map[string]string)
	response["message"] = "success"
	helper.ResponseJSON(w, http.StatusOK, response)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// delete existing token in cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	response := map[string]string{"message": "logout successful"}
	helper.ResponseJSON(w, http.StatusOK, response)
}
