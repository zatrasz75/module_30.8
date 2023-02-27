package postgres

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		constr string
	}
	var tests []struct {
		name    string
		args    args
		want    *Storage
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.constr)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_DelTaskLabel(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		taskID int
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			if err := s.DelTaskLabel(tt.args.taskID); (err != nil) != tt.wantErr {
				t.Errorf("DelTaskLabel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_DeleteLabel(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		labelId int
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			if err := s.DeleteLabel(tt.args.labelId); (err != nil) != tt.wantErr {
				t.Errorf("DeleteLabel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_DeleteTask(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		id int
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			if err := s.DeleteTask(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_DeleteUser(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		id int
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			if err := s.DeleteUser(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_LabelTask(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		lt Tasks_labels
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			if err := s.LabelTask(tt.args.lt); (err != nil) != tt.wantErr {
				t.Errorf("LabelTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_NameLabels(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		name string
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			got, err := s.NameLabels(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("NameLabels() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NameLabels() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_NewLabel(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		l Labels
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			got, err := s.NewLabel(tt.args.l)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewLabel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewLabel() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_NewTask(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		t Task
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			got, err := s.NewTask(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewTask() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_NewUser(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		u Users
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			got, err := s.NewUser(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_Tasks(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		taskID   int
		authorID int
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    []Task
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			got, err := s.Tasks(tt.args.taskID, tt.args.authorID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Tasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tasks() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_TasksAuthor(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		author string
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    []Task
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			got, err := s.TasksAuthor(tt.args.author)
			if (err != nil) != tt.wantErr {
				t.Errorf("TasksAuthor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TasksAuthor() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_TasksAuthorId(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		authorID int
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    []Task
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			got, err := s.TasksAuthorId(tt.args.authorID)
			if (err != nil) != tt.wantErr {
				t.Errorf("TasksAuthorId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TasksAuthorId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_TasksLabelId(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		labelId int
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    []Task
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			got, err := s.TasksLabelId(tt.args.labelId)
			if (err != nil) != tt.wantErr {
				t.Errorf("TasksLabelId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TasksLabelId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_UpdateTaskContent(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		id      int
		content string
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			if err := s.UpdateTaskContent(tt.args.id, tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("UpdateTaskContent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_UpdateTaskTitle(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		id    int
		title string
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			if err := s.UpdateTaskTitle(tt.args.id, tt.args.title); (err != nil) != tt.wantErr {
				t.Errorf("UpdateTaskTitle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
