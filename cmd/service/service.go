package main

import (
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

//var db storage.Interface

func main() {

	db, err := storage.New(connStr)
	if err != nil {
		elog.Fatal("Нет подключения к БД \n", err.Error())
	}

	//Автор
	userId, err := db.NewUser(storage.Users{
		Name: "Name",
	})
	if err != nil {
		fmt.Println(err)
	}
	ilog.Println(userId)

	//задача
	z := storage.Task{
		Title:   "Первая задача",
		Content: "Создавать новые задачи",
	}
	task, err := db.NewTask(z)
	if err != nil {
		elog.Println(err)
	}
	ilog.Printf("id новой задачи :%v\n", task)

	//----------------------------------------------------------------

	//Все задачи
	tasks, err := db.Tasks(0, 0)
	if err != nil {
		elog.Println(err)
	}
	ilog.Printf("Список задач: \n %v\n", tasks)

	//Задача по id
	aut, err := db.TasksAuthorId(1)
	if err != nil {
		elog.Println(err)
	}
	ilog.Printf("ищет задачи по идентификатору автора: \n%v\n", aut)

	// Задача по автору
	aut, err = db.TasksAuthor("Максим")
	if err != nil {
		elog.Println(err)
	}
	ilog.Printf("ищет задачи по имени автора: \n%v\n", aut)

	//ID метки по имени
	id, err := db.NameLabels("Метка 1") //==================================================
	if err != nil {
		elog.Println(err)
	}
	ilog.Printf("ищет ID по имени метки : \n%v\n", id)

	// Задачи по id метки
	aut, err = db.TasksLabelId(id)
	if err != nil {
		elog.Println(err)
	}
	ilog.Printf("ищет задачи по идентификатору метки : \n%v\n", aut)

	// Задачи по названию метки
	aut, err = db.TasksLabel("Метка 3")
	if err != nil {
		elog.Println(err)
	}
	ilog.Printf("ищет задачи по названию метки : \n%v\n", aut)

	// Обновляет заголовок задачи по id
	err = db.UpdateTaskTitle(4, "обновление")
	if err != nil {
		elog.Println(err)
	}
	// Обновляет текст задачи по id
	err = db.UpdateTaskContent(4, "прошло успешно")
	if err != nil {
		elog.Println(err)
	}
	// Удалить задачу
	err = db.DeleteTask(task)
	if err != nil {
		elog.Println(err)
	}
	// Удалить пользователя
	err = db.DeleteTask(userId)
	if err != nil {
		elog.Println(err)
	}

}
