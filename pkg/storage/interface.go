package storage

import "DB_Apps/pkg/postgres"

type Interface interface {
	Tasks(int, int) ([]postgres.Task, error)
	NewTask(postgres.Task) (int, error)
}
