package studentcontroller

import (
	"bytes"
	"encoding/json"
	htemplate "html/template"
	"net/http"
	"strconv"
	"text/template"

	"github.com/pjw1702/go-crud-capital/entities"
	"github.com/pjw1702/go-crud-capital/models/studentmodel"
)

var studentModel = studentmodel.New()

// show web page updated index.html(add, edit, delete) reflect through 'GetData'
func Index(w http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{
		"data": htemplate.HTML(GetData()),
	}

	temp, _ := template.ParseFiles("views/student/index.html")
	temp.Execute(w, data)
}

// Get data from web page table
func GetData() string {

	buffer := &bytes.Buffer{}

	temp, _ := template.New("data.html").Funcs(template.FuncMap{
		"increment": func(a, b int) int {
			return a + b
		},
	}).ParseFiles("views/student/data.html")

	var student []entities.Student
	err := studentModel.FindAll(&student)
	if err != nil {
		panic(err)
	}

	// fmt.Println(student)

	data := map[string]interface{}{
		"student": student,
	}

	temp.ExecuteTemplate(buffer, "data.html", data)

	return buffer.String()

}

// Get submit form to insert or update
func GetForm(w http.ResponseWriter, r *http.Request) {

	queryString := r.URL.Query()
	id, err := strconv.ParseInt(queryString.Get("id"), 10, 64)

	var data map[string]interface{}

	// insert form
	if err != nil {
		data = map[string]interface{}{
			"title": "Add student data",
		}
		// update form
	} else {
		var student entities.Student
		// Query to student DB (Read)
		err := studentModel.Find(id, &student)
		if err != nil {
			panic(err)
		}

		data = map[string]interface{}{
			"title":   "Edit student data",
			"student": student,
		}
	}

	temp, _ := template.ParseFiles("views/student/form.html")
	temp.Execute(w, data)
}

// Create, Update
func Store(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		r.ParseForm()
		var student entities.Student

		student.Name = r.Form.Get("full_name")
		student.Gender = r.Form.Get("gender")
		student.Birthplace = r.Form.Get("birthplace")
		student.Birthday = r.Form.Get("birthday")
		student.Address = r.Form.Get("address")

		id, err := strconv.ParseInt(r.Form.Get("id"), 10, 64)
		var data map[string]interface{}
		if err != nil {
			// insert data
			err := studentModel.Create(&student)
			if err != nil {
				ResponseError(w, http.StatusInternalServerError, err.Error())
				return
			}
			data = map[string]interface{}{
				"messgae": "data changed successfully",
				"data":    htemplate.HTML(GetData()),
			}
		} else {
			// update data
			student.Id = id
			err := studentModel.Update(student)
			if err != nil {
				ResponseError(w, http.StatusInternalServerError, err.Error())
				return
			}
			data = map[string]interface{}{
				"messgae": "data changed successfully",
				"data":    htemplate.HTML(GetData()),
			}
		}

		ResponseJson(w, http.StatusOK, data)
	}
}

// Delete
func Delete(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	id, err := strconv.ParseInt(r.Form.Get("id"), 10, 64)
	if err != nil {
		panic(err)
	}
	err = studentModel.Delete(id)
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"message": "Data deleted successfully",
		"data":    htemplate.HTML(GetData()),
	}
	ResponseJson(w, http.StatusOK, data)

}

func ResponseError(w http.ResponseWriter, code int, message string) {
	ResponseJson(w, code, map[string]string{"error": message})
}

func ResponseJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "applicaion/json")
	w.WriteHeader(code)
	w.Write(response)
}
