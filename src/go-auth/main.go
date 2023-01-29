package main

import (
	"fmt"
	"net/http"

	authcontroller "github.com/pjw1702/go-auth/controllers"
)

func main() {

	http.HandleFunc("/", authcontroller.Index)
	http.HandleFunc("/login", authcontroller.Login)
	http.HandleFunc("/logout", authcontroller.Logout)
	http.HandleFunc("/register", authcontroller.Register)

	fmt.Println("Server walk in: http://localhost:3000")
	http.ListenAndServe(":3000", nil)

}
