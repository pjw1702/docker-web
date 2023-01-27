package patientcontroller

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/pjw1702/go-crud-form-validation/entities"
	"github.com/pjw1702/go-crud-form-validation/libraries"
	"github.com/pjw1702/go-crud-form-validation/models"
)

// create validator
var validation = libraries.NewValidation()
var patientModel = models.NewPatientModel()

// Get main page: index.html
func Index(response http.ResponseWriter, request *http.Request) {

	// get all informations of 'patient' DB
	patient, _ := patientModel.FindAll()

	// show web page-index.html's table by getting value 'patient'
	data := map[string]interface{}{
		"patient": patient,
	}

	temp, err := template.ParseFiles("views/patient/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)
}

// Get submit form: add.html and add data
func Add(response http.ResponseWriter, request *http.Request) {

	// Get submit form: add.html
	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/patient/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {

		// Query submit form
		request.ParseForm()

		// Post
		var patient entities.Patient
		patient.Full_Name = request.Form.Get("full_name")
		patient.Number_Patient = request.Form.Get("patient_number")
		patient.Gender = request.Form.Get("gender")
		patient.Birthplace = request.Form.Get("birthplace")
		patient.Birthday = request.Form.Get("birthday")
		patient.Address = request.Form.Get("address")
		patient.Number_HP = request.Form.Get("phone_number")

		var data = make(map[string]interface{})

		// get invalidate message by struct of 'patient'
		vErrors := validation.Struct(patient)

		// check validate
		if vErrors != nil {
			data["patient"] = patient
			data["validation"] = vErrors
		} else {
			data["message"] = "Patient data saved successfully"
			patientModel.Create(patient)
		}

		//patientModel.Create(patient)
		//data := map[string]interface{}{
		//	"message": "Patient data saved successfully",
		//}

		temp, _ := template.ParseFiles("views/patient/add.html")
		temp.Execute(response, data)
	}

}

func Edit(response http.ResponseWriter, request *http.Request) {

	// Get submit form: edit.html
	if request.Method == http.MethodGet {

		// Parsing an specific id and check matching
		queryString := request.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

		var patient entities.Patient
		patientModel.Find(id, &patient)

		data := map[string]interface{}{
			"patient": patient,
		}

		temp, err := template.ParseFiles("views/patient/edit.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, data)
	} else if request.Method == http.MethodPost {

		// Query submit form
		request.ParseForm()

		// Post
		var patient entities.Patient
		patient.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
		patient.Full_Name = request.Form.Get("full_name")
		patient.Number_Patient = request.Form.Get("patient_number")
		patient.Gender = request.Form.Get("gender")
		patient.Birthplace = request.Form.Get("birthplace")
		patient.Birthday = request.Form.Get("birthday")
		patient.Address = request.Form.Get("address")
		patient.Number_HP = request.Form.Get("phone_number")

		var data = make(map[string]interface{})

		// get invalidate message by struct of 'patient'
		vErrors := validation.Struct(patient)

		// check validate
		if vErrors != nil {
			data["patient"] = patient
			data["validation"] = vErrors
		} else {
			data["message"] = "Patient data successfully updated"
			patientModel.Update(patient)
		}

		//patientModel.Create(patient)
		//data := map[string]interface{}{
		//	"message": "Patient data saved successfully",
		//}

		temp, _ := template.ParseFiles("views/patient/edit.html")
		temp.Execute(response, data)
	}

}

func Delete(response http.ResponseWriter, request *http.Request) {

	// Parsing an specific id and check matching
	queryString := request.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	patientModel.Delete(id)

	// Redirect automatically after task of delete
	http.Redirect(response, request, "/patient", http.StatusSeeOther)

}
