// go get github.com/go-sql-driver/mysql github.com/go-playground/validator/v10
package main

import (
	"net/http"

	"github.com/pjw1702/go-crud-form-validation/controllers/patientcontroller"
)

func main() {

	http.HandleFunc("/", patientcontroller.Index)
	http.HandleFunc("/patient", patientcontroller.Index)
	http.HandleFunc("/patient/index", patientcontroller.Index)
	http.HandleFunc("/patient/add", patientcontroller.Add)
	http.HandleFunc("/patient/edit", patientcontroller.Edit)
	http.HandleFunc("/patient/delete", patientcontroller.Delete)

	http.ListenAndServe(":3000", nil)
}
