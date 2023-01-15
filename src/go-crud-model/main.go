package main

import (
	"net/http"

	"github.com/pjw1702/go-crud-capital/controllers/studentcontroller"
)

func main() {
	http.HandleFunc("/", studentcontroller.Index)
	http.HandleFunc("/student/get_form", studentcontroller.GetForm)
	http.HandleFunc("/student/store", studentcontroller.Store)
	http.HandleFunc("/student/delete", studentcontroller.Delete)

	http.ListenAndServe(":8000", nil)
}
