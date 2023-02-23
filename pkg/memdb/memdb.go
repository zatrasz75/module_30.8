package memdb

import "DB_Apps/pkg/postgres"

type DB []postgres.Task

func (db DB) Tasks(int, int) ([]postgres.Task, error) {
	return db, nil
}

func (db DB) NewTasks(t postgres.Task) (int, error) {
	return 0, nil
}
