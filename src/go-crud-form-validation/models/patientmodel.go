package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/pjw1702/go-crud-form-validation/config"
	"github.com/pjw1702/go-crud-form-validation/entities"
)

// DB object mapping(ORM) module set
type PatientModel struct {
	conn *sql.DB
}

// Connect to DB
func NewPatientModel() *PatientModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &PatientModel{
		conn: conn,
	}
}

// Read all DB and store in slice called 'patient'
func (p *PatientModel) FindAll() ([]entities.Patient, error) {
	rows, err := p.conn.Query("select * from patient")
	if err != nil {
		return []entities.Patient{}, err
	}
	defer rows.Close()

	var dataPatient []entities.Patient
	for rows.Next() {
		var patient entities.Patient
		rows.Scan(&patient.Id,
			&patient.Full_Name,
			&patient.Number_Patient,
			&patient.Gender,
			&patient.Birthplace,
			&patient.Birthday,
			&patient.Address,
			&patient.Number_HP)

		// convert webpage table's field of Gender:number 1 or 2 to men or women
		if patient.Gender == "1" {
			patient.Gender = "Men"
		} else {
			patient.Gender = "Women"
		}

		// 2006-01-02 => yyyy-mm-dd
		birthday, _ := time.Parse("2006-01-02", patient.Birthday)
		// 02-01-2006 => dd-mm-yyyy
		patient.Birthday = birthday.Format("02-01-2006")

		dataPatient = append(dataPatient, patient)
	}

	return dataPatient, nil
}

// Insert data to DB
func (p *PatientModel) Create(patient entities.Patient) bool {

	result, err := p.conn.Exec("insert into patient (full_name, number_patient, gender, birthplace, birthday, address, number_hp) values(?,?,?,?,?,?,?)",
		patient.Full_Name, patient.Number_Patient, patient.Gender, patient.Birthplace, patient.Birthday, patient.Address, patient.Number_HP)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}

// Read DB information Only parts that match a specific ID
func (p *PatientModel) Find(id int64, patient *entities.Patient) error {

	return p.conn.QueryRow("select * from patient where id = ?", id).Scan(
		&patient.Id,
		&patient.Full_Name,
		&patient.Number_Patient,
		&patient.Gender,
		&patient.Birthplace,
		&patient.Birthday,
		&patient.Address,
		&patient.Number_HP)

}

// Update DB
func (p *PatientModel) Update(patient entities.Patient) error {

	_, err := p.conn.Exec(
		"update patient set full_name = ?, number_patient = ?, gender = ?, birthplace = ?, birthplace = ?, address = ?, number_hp = ? where id = ?",
		patient.Full_Name, patient.Number_Patient, patient.Gender, patient.Birthplace, patient.Birthday, patient.Address, patient.Number_HP, patient.Id)

	if err != nil {
		return err
	}

	return nil

}

// Delete DB
func (p *PatientModel) Delete(id int64) {
	p.conn.Exec("delete from patient where id = ?", id)
}
