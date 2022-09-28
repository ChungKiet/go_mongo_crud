package student_storage

import "go-crud-mongo/student_model"

type StudentStorageService interface {
	CreateStudent(data *student_model.Student) (*student_model.Student, error)
	GetStudent(sID string) (*student_model.Student, error)
	GetAllStudent() ([]*student_model.Student, error)
	UpdateStudent(sID string, Name string, Class int, Gender string) error
	DeleteStudent(sID string) error
}
