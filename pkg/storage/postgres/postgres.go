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

// Task Задачи.
type Task struct {
	ID         int
	Opened     int64
	Closed     int64
	AuthorID   int
	AssignedID int
	Title      string
	Content    string
}

// Users Пользователи
type Users struct {
	ID   int
	Name string
}

// Labels Метки
type Labels struct {
	ID   int
	Name string
}

// Tasks_labels Метки для задач
type Tasks_labels struct {
	Task_id  int
	Label_id int
}

// NewUser Создает пользователя
func (s *Storage) NewUser(u Users) (int, error) {
	var id int
	err := s.db.QueryRow(context.Background(), `
		INSERT INTO users (name)
		VALUES ($1) RETURNING id;
		`,
		u.Name,
	).Scan(&id)

	return id, err
}

// NewTask создаёт новую задачу и возвращает её id.
func (s *Storage) NewTask(t Task) (int, error) {
	var id int
	err := s.db.QueryRow(context.Background(), `
		INSERT INTO tasks (author_id, assigned_id, title, content)
		VALUES ($1, $2, $3, $4) RETURNING id;
		`,
		t.AuthorID,
		t.AssignedID,
		t.Title,
		t.Content,
	).Scan(&id)

	return id, err
}

// NewLabel Создает новую метку и возвращает её id.
func (s *Storage) NewLabel(l Labels) (int, error) {
	var id int
	err := s.db.QueryRow(context.Background(), `
		INSERT INTO labels (name)
		VALUES ($1) RETURNING id;
		`,
		l.Name,
	).Scan(&id)

	return id, err
}

// LabelTask Отмечаем задачу меткой
func (s *Storage) LabelTask(lt Tasks_labels) error {
	_, err := s.db.Exec(context.Background(), `
		INSERT INTO tasks_labels (task_id, label_id)
		VALUES ($1, $2) ;
		`,
		lt.Task_id,
		lt.Label_id,
	)
	if err != nil {
		return err
	}
	return err
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

// TasksAuthorId ищет задачи по идентификатору автора
func (s *Storage) TasksAuthorId(authorID int) ([]Task, error) {
	var tasks []Task
	query := `
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
			$1 = 0 OR author_id = $1
		ORDER BY id;`
	rows, err := s.db.Query(context.Background(), query, authorID)

	if err != nil {
		return nil, err
	}

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
		tasks = append(tasks, t)

	}
	return tasks, rows.Err()
}

// TasksAuthor ищет задачи по имени автора
func (s *Storage) TasksAuthor(author string) ([]Task, error) {
	var tasks []Task
	query := `
		SELECT 
			id,
			opened,
			closed,
			author_id,
			assigned_id,
			title,
			content
		FROM tasks
		WHERE author_id in (SELECT id FROM users WHERE name=$1)
		ORDER BY id;`
	rows, err := s.db.Query(context.Background(), query, author)

	if err != nil {
		return nil, err
	}

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
		tasks = append(tasks, t)

	}
	return tasks, rows.Err()
}

// NameLabels Возвращает ID метки по ее имени
func (s *Storage) NameLabels(name string) (int, error) {

	query := `
		SELECT 
			id,
			name
		FROM labels
		WHERE
			$1 = name OR name = $1
		ORDER BY id;`

	rows, err := s.db.Query(context.Background(), query, name)
	if err != nil {
		return 0, err
	}
	var id int
	for rows.Next() {
		var t Labels
		err = rows.Scan(
			&t.ID,
			&t.Name,
		)
		if err != nil {
			return 0, err
		}
		id = t.ID
	}
	return id, rows.Err()
}

// TasksLabelId ищет задачи по идентификатору метки
func (s *Storage) TasksLabelId(labelId int) ([]Task, error) {
	query := `SELECT 
					id,
					opened,
					closed,
					author_id,
					assigned_id,
					title,
					content
              FROM tasks
              WHERE id in (SELECT task_id FROM tasks_labels WHERE label_id=$1);`

	rows, err := s.db.Query(context.Background(), query, labelId)

	if err != nil {
		return nil, err
	}

	var tasks []Task

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
		tasks = append(tasks, t)

	}
	return tasks, rows.Err()
}

// UpdateTaskTitle обновляет заголовок задачи по id
func (s *Storage) UpdateTaskTitle(id int, title string) error {
	query := `UPDATE tasks SET title = $2 WHERE id = $1;`
	_, err := s.db.Exec(context.Background(), query, id, title)
	if err != nil {
		return err
	}
	return nil
}

// UpdateTaskContent обновляет текст задачи по id
func (s *Storage) UpdateTaskContent(id int, content string) error {
	query := `UPDATE tasks SET content = $2 WHERE id = $1;`
	_, err := s.db.Exec(context.Background(), query, id, content)
	if err != nil {
		return err
	}
	return nil
}

// DelTaskLabel удаляет пометку задачи по id
func (s *Storage) DelTaskLabel(taskID int) error {
	delet := `
		DELETE FROM tasks_labels WHERE task_id = $1;
	`
	_, err := s.db.Exec(context.Background(), delet, taskID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTask удаляет задачу по id
func (s *Storage) DeleteTask(id int) error {
	delet := `
		DELETE FROM tasks WHERE id = $1;
	`
	_, err := s.db.Exec(context.Background(), delet, id)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser удаляет пользователя по id
func (s *Storage) DeleteUser(id int) error {
	delet := `
		DELETE FROM users WHERE id = $1;
	`
	_, err := s.db.Exec(context.Background(), delet, id)
	if err != nil {
		return err
	}
	return nil
}

// DeleteLabel удаляет метку по id
func (s *Storage) DeleteLabel(labelId int) error {
	delet := `
		DELETE FROM labels WHERE id = $1;
	`
	_, err := s.db.Exec(context.Background(), delet, labelId)
	if err != nil {
		return err
	}
	return nil
}
