package storage

// Task Задача.
type Task struct {
	ID         int
	Opened     int64
	Closed     int64
	AuthorID   int
	AssignedID int
	Title      string
	Content    string
}

// Labels Метки
type Labels struct {
	ID   int
	Name string
}

// Tasks_labels Метки задач
type Tasks_labels struct {
	task_id  int
	label_id int
}

// Users Пользователи
type Users struct {
	ID   int
	Name string
}
