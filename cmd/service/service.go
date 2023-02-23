package main

import (
	"DB_Apps/pkg/postgres"
	"DB_Apps/pkg/storage"
	"fmt"
	"log"
	"os"
)

const (
	Host     = "localhost"
	Port     = 5432
	Users    = "postgres"
	Password = "rootroot"
	Dbname   = "tasks" //
)

var elog = log.New(os.Stderr, "service error\t", log.Ldate|log.Ltime|log.Lshortfile)
var ilog = log.New(os.Stdout, "service info\t", log.Ldate|log.Ltime)

var (
	// Подключение к БД
	connStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai", Host, Port, Users, Password, Dbname)
)
var db storage.Interface

func main() {

	db, err := postgres.New(connStr)
	if err != nil {
		elog.Fatal("Нет подключения к БД \n", err.Error())
	}
	//db = memdb.DB{}
	tasks, err := db.Tasks(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	ilog.Println(tasks)

	//id, _ := db.NewTask(postgres.Task{
	//	ID:         5,
	//	Title:      "memdb task 5",
	//	Content:    "new task 5",
	//	AuthorID:   5,
	//	AssignedID: 5,
	//})
	//ilog.Println(id)
	//
	//id, _ = db.NewTask(postgres.Task{
	//	ID:         6,
	//	Title:      "memdb task 6",
	//	Content:    "new task 6",
	//	AuthorID:   6,
	//	AssignedID: 6,
	//})
	//ilog.Println(id)

	//============================================================

	//s := postgres.Task{
	//	Title:   "ups",
	//	Content: "Проверка связи",
	//}
	//
	//id, err = db.NewTask(s)
	//if err != nil {
	//	return
	//}
	//fmt.Println(id)

	//err = db.TaskDel(1)
	//if err != nil {
	//	return
	//}

	//db.Tasks(task, )

}
