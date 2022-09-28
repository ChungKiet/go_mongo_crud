package student_model

import (
	"context"
)

type Student struct {
	SID    string `json:"s_id" bson:"s_id"`
	Name   string `json:"name" bson:"name" binding:"required"`
	Class  int    `json:"class" bson:"class" binding:"required"`
	Gender string `json:"gender" bson:"gender" binding:"required,oneof=MALE, FEMALE"`
}

type StudentMethod interface {
	CreateStudent(ctx context.Context, Name string, Class int, Gender string) (Student, error)
	GetStudent(ctx context.Context, sID string) (Student, error)
	GetAllStudent(ctx context.Context) ([]Student, error)
	UpdateStudent(ctx context.Context, sID string, Name string, Class int, Gender string) (Student, error)
	DeleteStudent(ctx context.Context, sID string) (Student, error)
}
