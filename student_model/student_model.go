package student_model

import (
	"context"
)

type Student struct {
	SID    string `json:"s_id" bson:"s_id" binding:"required"`
	Name   string `json:"name" bson:"name" binding:"required"`
	Class  int    `json:"class" bson:"class" binding:"required"`
	Gender string `json:"gender" bson:"gender" binding:"required"`
}

type StudentMethod interface {
	CreateStudent(ctx context.Context, Name string, Class int, Gender string) (Student, error)
	GetStudent(ctx context.Context, sID string) (Student, error)
	GetAllStudent(ctx context.Context) ([]Student, error)
	UpdateStudent(ctx context.Context, sID string, Name string, Class int, Gender string) (Student, error)
	DeleteStudent(ctx context.Context, sID string) (Student, error)
}

//var wg sync.WaitGroup
//
//
//func insertData(thread int){
//	fmt.Println(thread);
//
//	db, err := sql.Open("mysql", "root:root@tcp(172.22.0.2:3306)/crm_mediatama")
//
//	if err != nil {
//		panic(err)
//	}
//
//	db.SetConnMaxLifetime(time.Minute * 30)
//	db.SetMaxOpenConns(10)
//	db.SetMaxIdleConns(10)
//
//	fmt.Println("Mysql Connnection Ready")
//
//	sqlStr := "INSERT INTO visits(date, school_id, user_id, discription) VALUES "
//	vals := []interface{}{}
//
//	for i := 0;i < 10;i++ {
//		sqlStr += "(?, ?, ?, ?),"
//		vals = append(vals, "2020-01-01", "1", "1","des-golang-" + strconv.Itoa(i))
//	}
//
//	// trim the last ,
//	sqlStr = sqlStr[0:len(sqlStr)-1]
//
//	// prepare the statement
//	stmt,errStmt := db.Prepare(sqlStr)
//
//
//	if(errStmt != nil){
//		panic(errStmt)
//	}
//
//	// format all vals at once
//	res,errExc := stmt.Exec(vals...)
//
//	if(errExc != nil){
//		panic(errExc)
//	}
//
//	fmt.Println(res.RowsAffected());
//
//	db.Close()
//
//	wg.Done();
//}
//
//func main() {
//	for i := 0; i < 5; i++ {
//		wg.Add(1)
//		go insertData(i)
//	}
//
//	wg.Wait()
//}
