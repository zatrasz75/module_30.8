package storage

import (
	"DB_Apps/pkg/storage/postgres"
)

// Interface Интерфейс БД.
// Этот интерфейс позволяет абстрагироваться от конкретной СУБД.
// Можно создать реализацию БД в памяти для модульных тестов.
type Interface interface {
	NewUser(postgres.Users) (int, error)
	NewTask(postgres.Task) (int, error)
	NewLabel(postgres.Labels) (int, error)
	LabelTask(lt postgres.Tasks_labels) error
	Tasks(int, int) ([]postgres.Task, error)
	TasksAuthor(string) ([]postgres.Task, error)
	TasksAuthorId(int) ([]postgres.Task, error)
	NameLabels(string) (int, error)
	TasksLabelId(int) ([]postgres.Task, error)
	UpdateTaskTitle(int, string) error
	UpdateTaskContent(int, string) error
	DelTaskLabel(int) error
	DeleteTask(int) error
	DeleteUser(int) error
	DeleteLabel(int) error
}
