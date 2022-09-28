package student_storage

import (
	"context"
	"errors"
	"go-crud-mongo/student_model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type StudentCollections struct {
	studentCollection *mongo.Collection
	ctx               context.Context
}

func NewStudentStorage(studentCollection *mongo.Collection, ctx context.Context) *StudentCollections {
	return &StudentCollections{
		studentCollection: studentCollection,
		ctx:               ctx,
	}
}

func (stuCol *StudentCollections) CreateStudent(data *student_model.Student) (*student_model.Student, error) {
	_, err := stuCol.studentCollection.InsertOne(stuCol.ctx, data)

	return data, err
}

func (stuCol *StudentCollections) GetStudent(sID string) (*student_model.Student, error) {
	data := &student_model.Student{}

	query := bson.D{bson.E{Key: "s_id", Value: sID}}
	err := stuCol.studentCollection.FindOne(stuCol.ctx, query).Decode(&data)

	if err != nil {
		return nil, errors.New("Cannot get config")
	}

	return data, nil
}

func (stuCol *StudentCollections) GetAllStudent() ([]*student_model.Student, error) {
	var allStudents []*student_model.Student

	cursor, err := stuCol.studentCollection.Find(stuCol.ctx, bson.D{{}})

	if err != nil {
		return nil, errors.New("Cannot get config")
	}

	for cursor.Next(stuCol.ctx) {
		var stud student_model.Student
		err := cursor.Decode(&stud)
		if err != nil {
			return nil, err
		}
		// Append config in map
		allStudents = append(allStudents, &stud)
	}

	err = cursor.Close(stuCol.ctx)

	if err != nil {
		return nil, err
	}

	// If there are no config, then return nil
	if len(allStudents) == 0 {
		return nil, errors.New("no documents found")
	}

	return allStudents, nil
}

func (stuCol *StudentCollections) UpdateStudent(sID string, Name string, Class int, Gender string) error {
	filter := bson.M{"s_id": sID}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "name", Value: Name}, primitive.E{Key: "class", Value: Class}, primitive.E{Key: "gender", Value: Gender}}}}
	_, err := stuCol.studentCollection.UpdateOne(stuCol.ctx, filter, update)
	return err
}

func (stuCol *StudentCollections) DeleteStudent(sID string) error {
	filter := bson.M{"s_id": sID}
	_, err := stuCol.studentCollection.DeleteOne(stuCol.ctx, filter)
	return err
}
