package main

import (
	"DB_Apps/pkg/storage"
	"DB_Apps/pkg/storage/postgres"
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
	//	Заглушка для тестов
	//db = memdb.DB{}

	fmt.Println("Создавать новые задачи -------------------------------------------------------")

	//Создаем нового автора
	userId, err := db.NewUser(postgres.Users{
		Name: "Михаил",
	})
	if err != nil {
		elog.Println(err)
	}
	ilog.Printf("id нового автора :\n%v\n", userId)

	// Создаём новую задачу для нового автора
	taskId, err := db.NewTask(postgres.Task{
		AuthorID:   userId,
		AssignedID: userId,
		Title:      "Задача",
		Content:    "Krex Pex Fex",
	})
	if err != nil {
		elog.Println(err)
	}
	ilog.Printf("id новой задачи для нового автора:\n%v\n", taskId)

	// Создаем новую метку
	labelId, err := db.NewLabel(postgres.Labels{
		Name: "Метка",
	})
	if err != nil {
		elog.Println(err)
	}
	ilog.Printf("id новой метки :\n%v\n", labelId)

	// Отмечаем задачу меткой
	err = db.LabelTask(postgres.Tasks_labels{
		Task_id:  taskId,
		Label_id: labelId,
	})
	if err != nil {
		elog.Println(err)
	}

	fmt.Println("Получать список всех задач ----------------------------------------------------")

	// Получать список всех задач
	tasks, err := db.Tasks(0, 0)
	if err != nil {
		elog.Println(err)
	}
	ilog.Printf("Список всех задач: \n %v\n", tasks)

	fmt.Println("Получать список задач по автору -----------------------------------------------")

	// Ищет задачи по имени автора
	aut, err := db.TasksAuthor("Михаил")
	if err != nil {
		elog.Println(err)
	}
	ilog.Printf("Ищет задачи по имени автора: \n%v\n", aut)

	//Ищет задачи по id автора
	aut, err = db.TasksAuthorId(userId)
	if err != nil {
		elog.Println(err)
	}
	ilog.Printf("ищет задачи по идентификатору автора: \n%v\n", aut)

	fmt.Println("Получать список задач по метке ------------------------------------------------")

	//Получить ID метки по её имени
	id, err := db.NameLabels("Метка")
	if err != nil {
		elog.Println(err)
	}
	ilog.Printf("Получить ID по её имени : \n%v\n", id)

	// Получить задачи по id метки
	aut, err = db.TasksLabelId(id)
	if err != nil {
		elog.Println(err)
	}
	ilog.Printf("ищет задачи по id метки : \n%v\n", aut)

	fmt.Println("Обновлять задачу по id -------------------------------------------------------")

	// Обновляет заголовок задачи по id
	err = db.UpdateTaskTitle(taskId, "обновление")
	if err != nil {
		elog.Println(err)
	}
	// Обновляет текст задачи по id
	err = db.UpdateTaskContent(taskId, "прошло успешно")
	if err != nil {
		elog.Println(err)
	}

	fmt.Println("Удалять задачу по id ---------------------------------------------------------")

	// Удаляет пометку задачи по id
	err = db.DelTaskLabel(taskId)
	if err != nil {
		elog.Println(err)
	}

	// Удалить задачу
	err = db.DeleteTask(taskId)
	if err != nil {
		elog.Println(err)
	}

	// Удалить автора
	err = db.DeleteUser(userId)
	if err != nil {
		elog.Println(err)
	}

	// Удалить метку
	err = db.DeleteLabel(labelId)
	if err != nil {
		elog.Println(err)
	}

}
