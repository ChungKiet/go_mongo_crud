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
	Gender string `json:"gender" bson:"gender" binding:"required"`
}

type GetStudentRequest struct {
	SID string `form:"s_id"`
}

type UpdateStudentRequest struct {
	SID    string `json:"s_id" bson:"s_id"`
	Name   string `json:"name" bson:"name" binding:"required"`
	Class  int    `json:"class" bson:"class" binding:"required"`
	Gender string `json:"gender" bson:"gender" binding:"required"`
}

type StudentController struct {
	StudentService student_storage.StudentStorageService
}

type DeleteStudentRequest struct {
	SID string `form:"s_id"`
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

func (s *StudentController) GetStudent(ctx *gin.Context) {

	stud := student_model.Student{}
	request := &GetStudentRequest{}
	if err := ctx.ShouldBindQuery(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, request)
		return
	}

	stud.SID = request.SID

	result, err := s.StudentService.GetStudent(stud.SID)

	if err != nil || result == nil {
		fmt.Print("Error get config")
		ctx.JSON(http.StatusOK, err)
	}

	ctx.JSON(http.StatusOK, result)
}

func (s *StudentController) GetAllStudent(ctx *gin.Context) {

	allStudent, err := s.StudentService.GetAllStudent()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, allStudent)
}

func (s *StudentController) UpdateStudent(ctx *gin.Context) {

	stud := student_model.Student{}
	request := &UpdateStudentRequest{}
	if err := ctx.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	stud.SID = request.SID
	stud.Name = request.Name
	stud.Class = request.Class
	stud.Gender = request.Gender

	err := s.StudentService.UpdateStudent(request.SID, request.Name, request.Class, request.Gender)

	if err != nil {
		fmt.Print("Error update config")
		ctx.JSON(http.StatusOK, err)
	}

	ctx.JSON(http.StatusOK, "Update Successful!")
}

func (s *StudentController) DeleteStudent(ctx *gin.Context) {

	stud := student_model.Student{}
	request := &DeleteStudentRequest{}
	if err := ctx.ShouldBindQuery(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	stud.SID = request.SID

	err := s.StudentService.DeleteStudent(request.SID)

	if err != nil {
		fmt.Print("Error update config")
		ctx.JSON(http.StatusOK, err)
	}

	ctx.JSON(http.StatusOK, "Delete Successful!")
}

func (s *StudentController) RegisterUserRoutes(rg *gin.RouterGroup) {
	studentRoute := rg.Group("/student")
	studentRoute.POST("/create-student", s.CreateStudent)
	studentRoute.GET("/", s.GetAllStudent)
	studentRoute.GET("/get-student", s.GetStudent)
	studentRoute.PUT("/update-student", s.UpdateStudent)
	studentRoute.DELETE("/delete-student", s.DeleteStudent)
}
