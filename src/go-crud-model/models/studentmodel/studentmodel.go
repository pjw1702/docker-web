package studentmodel

import (
	"database/sql"

	"github.com/pjw1702/go-crud-capital/config"
	"github.com/pjw1702/go-crud-capital/entities"
)

// Student DB information
type Studentmodel struct {
	db *sql.DB
}

// Connect to student DB
func New() *Studentmodel {
	db, err := config.DBConnection()
	if err != nil {
		panic(err)
	}
	return &Studentmodel{db: db}
}

// Read student DB(all data)
// Find student DB information to display existing or updated information on a webpage
func (m *Studentmodel) FindAll(student *[]entities.Student) error {
	rows, err := m.db.Query("select * from student")
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var data entities.Student
		rows.Scan(
			&data.Id,
			&data.Name,
			&data.Gender,
			&data.Birthplace,
			&data.Birthday,
			&data.Address)

		*student = append(*student, data)
	}

	return nil
}

// Create student DB
// Add new information for student DB from submit form
func (m *Studentmodel) Create(student *entities.Student) error {
	result, err := m.db.Exec("insert into student (name, gender, birthplace, birthday, address) values(?,?,?,?,?)",
		student.Name, student.Gender, student.Birthplace, student.Birthday, student.Address)

	if err != nil {
		return err
	}

	// Get id through function of db.Exec
	lastInsertId, _ := result.LastInsertId()
	// You must set default value for "id" to "Auto Increment" on DBMS
	student.Id = lastInsertId
	return nil
}

// Read student DB(specific data)
// Find student DB information to edit information
func (m *Studentmodel) Find(id int64, student *entities.Student) error {
	return m.db.QueryRow("select * from student where id = ?", id).Scan(
		&student.Id,
		&student.Name,
		&student.Gender,
		&student.Birthplace,
		&student.Birthday,
		&student.Address)
}

// Update student DB
func (m *Studentmodel) Update(student entities.Student) error {

	_, err := m.db.Exec("update student set name = ?, gender = ?, birthplace = ?, birthday = ?, address = ? where id = ?",
		student.Name, student.Gender, student.Birthplace, student.Birthday, student.Address, student.Id)

	if err != nil {
		return err
	}

	return nil

}

// Delete student DB
func (m *Studentmodel) Delete(id int64) error {
	_, err := m.db.Exec("delete from student where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
