package memdb

import (
	"DB_Apps/pkg/storage/postgres"
)

// DB Пользовательский тип данных - реализация БД в памяти.
// Т.н. "заглушка".
type DB []postgres.Task

// NewLabel Выполнение контракта интерфейса storage.Interface
func (db DB) NewLabel(labels postgres.Labels) (int, error) {
	return 0, nil
}

// LabelTask Выполнение контракта интерфейса storage.Interface
func (db DB) LabelTask(lt postgres.Tasks_labels) error {
	return nil
}

// DelTaskLabel Выполнение контракта интерфейса storage.Interface
func (db DB) DelTaskLabel(i int) error {
	return nil
}

// DeleteLabel Выполнение контракта интерфейса storage.Interface
func (db DB) DeleteLabel(i int) error {
	return nil
}

// NewUser Выполнение контракта интерфейса storage.Interface
func (db DB) NewUser(u postgres.Users) (int, error) {
	return 0, nil
}

// TasksAuthorId Выполнение контракта интерфейса storage.Interface
func (db DB) TasksAuthorId(i int) ([]postgres.Task, error) {
	return db, nil
}

// TasksAuthor Выполнение контракта интерфейса storage.Interface
func (db DB) TasksAuthor(s string) ([]postgres.Task, error) {
	return db, nil
}

// NameLabels Выполнение контракта интерфейса storage.Interface
func (db DB) NameLabels(s string) (int, error) {
	return 0, nil
}

// TasksLabelId Выполнение контракта интерфейса storage.Interface
func (db DB) TasksLabelId(i int) ([]postgres.Task, error) {
	return db, nil
}

// UpdateTaskTitle Выполнение контракта интерфейса storage.Interface
func (db DB) UpdateTaskTitle(i int, s string) error {
	return nil
}

// UpdateTaskContent Выполнение контракта интерфейса storage.Interface
func (db DB) UpdateTaskContent(i int, s string) error {
	return nil
}

// DeleteTask Выполнение контракта интерфейса storage.Interface
func (db DB) DeleteTask(id int) error {
	return nil
}

// DeleteUser Выполнение контракта интерфейса storage.Interface
func (db DB) DeleteUser(id int) error {
	return nil
}

// Tasks Выполнение контракта интерфейса storage.Interface
func (db DB) Tasks(int, int) ([]postgres.Task, error) {
	return db, nil
}

// NewTask Выполнение контракта интерфейса storage.Interface
func (db DB) NewTask(postgres.Task) (int, error) {
	return 0, nil
}
