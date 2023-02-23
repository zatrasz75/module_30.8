package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Storage Хранилище данных.
type Storage struct {
	db *pgxpool.Pool
}

// New Конструктор.
func New(constr string) (*Storage, error) {
	db, err := pgxpool.Connect(context.Background(), constr)
	if err != nil {
		return nil, err
	}
	s := Storage{
		db: db,
	}
	return &s, nil
}

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

// NewTask создаёт новую задачу и возвращает её id.
func (s *Storage) NewTask(t Task) (int, error) {
	var id int
	err := s.db.QueryRow(context.Background(), `
		INSERT INTO tasks (title, content)
		VALUES ($1, $2) RETURNING id;
		`,
		t.Title,
		t.Content,
	).Scan(&id)
	return id, err
}

// Tasks возвращает список задач из БД.
func (s *Storage) Tasks(taskID, authorID int) ([]Task, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT 
			id,
			opened,
			closed,
			author_id,
			assigned_id,
			title,
			content
		FROM tasks
		WHERE
			($1 = 0 OR id = $1) AND
			($2 = 0 OR author_id = $2)
		ORDER BY id;
	`,
		taskID,
		authorID,
	)
	if err != nil {
		return nil, err
	}
	var tasks []Task
	// итерированное по результату выполнения запроса
	// и сканирование каждой строки в переменную
	for rows.Next() {
		var t Task
		err = rows.Scan(
			&t.ID,
			&t.Opened,
			&t.Closed,
			&t.AuthorID,
			&t.AssignedID,
			&t.Title,
			&t.Content,
		)
		if err != nil {
			return nil, err
		}
		// добавление переменной в массив результатов
		tasks = append(tasks, t)

	}
	// ВАЖНО не забыть проверить rows.Err()
	return tasks, rows.Err()
}

// TasksAuth возвращает список задач по автору
func (s *Storage) TasksAuth(authorName string) ([]Task, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT 
		tasks.id,
		tasks.opened,
		tasks.closed,
		tasks.author_id,
		tasks.assigned_id,
		tasks.title,
		tasks.content
		FROM tasks, users
		WHERE
			users.name = $1
			tasks.author_id = users.id
		ORDER BY id;
	`,
		authorName,
	)
	if err != nil {
		return nil, err
	}
	var tasks []Task
	// итерирование по результату выполнения запроса
	// и сканирование каждой строки в переменную
	for rows.Next() {
		var t Task
		err = rows.Scan(
			&t.ID,
			&t.Opened,
			&t.Closed,
			&t.AuthorID,
			&t.AssignedID,
			&t.Title,
			&t.Content,
		)
		if err != nil {
			return nil, err
		}
		// добавление переменной в массив результатов
		tasks = append(tasks, t)

	}
	// ВАЖНО не забыть проверить rows.Err()
	return tasks, rows.Err()
}

// TasksLbl возвращает список задач по метке
func (s *Storage) TasksLbl(label string) ([]Task, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT 
		tasks.id,
		tasks.opened,
		tasks.closed,
		tasks.author_id,
		tasks.assigned_id,
		tasks.title,
		tasks.content,
		FROM tasks, labels, tasks_labels
		WHERE
			lables.name = $1
			tasks_labels.label_id = labels.id
			tasks.id = tasks_labels.task_id
		ORDER BY id;
	`,
		label,
	)
	if err != nil {
		return nil, err
	}
	var tasks []Task
	// итерированное по результату выполнения запроса
	// и сканирование каждой строки в переменную
	for rows.Next() {
		var t Task
		err = rows.Scan(
			&t.ID,
			&t.Opened,
			&t.Closed,
			&t.AuthorID,
			&t.AssignedID,
			&t.Title,
			&t.Content,
		)
		if err != nil {
			return nil, err
		}
		// добавление переменной в массив результатов
		tasks = append(tasks, t)

	}
	// ВАЖНО не забыть проверить rows.Err()
	return tasks, rows.Err()
}

// TaskUpd обновляет задачу по id
func (s *Storage) TaskUpd(taskID int, attribute string, target interface{}) (Task, error) {
	var updatedTask Task

	err := s.db.QueryRow(context.Background(), `
	UPDATE tasks
	SET $2 = '$3'
	WHERE id = $1
	`,
		taskID,
		attribute,
		target,
	).Scan(&updatedTask)

	return updatedTask, err
}

// TaskDel удаляет задачу по id
func (s *Storage) TaskDel(taskID int) error {
	err := s.db.QueryRow(context.Background(), `
		DELETE FROM tasks
		WHERE id = $1;
		`,
		taskID,
	).Scan()
	return err
}
