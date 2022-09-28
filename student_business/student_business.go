package student_business

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-crud-mongo/student_model"
	"go-crud-mongo/student_storage"
	"net/http"
)

type CreateStudentRequest struct {
	SID    string `json:"s_id" bson:"s_id"`
	Name   string `json:"name" bson:"name" binding:"required"`
	Class  int    `json:"class" bson:"class" binding:"required"`
	Gender string `json:"gender" bson:"gender" binding:"required,oneof=MALE, FEMALE"`
}

type StudentController struct {
	StudentService student_storage.StudentStorageService
}

func New(studentService student_storage.StudentStorageService) StudentController {
	return StudentController{
		StudentService: studentService,
	}
}

func (s *StudentController) CreateStudent(ctx *gin.Context) {

	stud := student_model.Student{}
	request := &CreateStudentRequest{}
	if err := ctx.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	stud.SID = request.SID
	stud.Name = request.Name
	stud.Class = request.Class
	stud.Gender = request.Gender

	result, err := s.StudentService.CreateStudent(&stud)

	if err != nil || result == nil {
		fmt.Print("Error create config")
		ctx.JSON(http.StatusOK, err)
	}

	ctx.JSON(http.StatusOK, result)
}
