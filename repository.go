package main

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type TaskRepository struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) *TaskRepository {
	db.MustExec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			tag TEXT NOT NULL,
			description TEXT NOT NULL,
			start_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			end_at DATETIME DEFAULT NULL
		)
	`)
	return &TaskRepository{db: db}
}

func (t TaskRepository) FindByStartDate(startAt time.Time) (*[]*Task, error) {

	tasks := &[]*Task{}

	if err := t.db.Select(tasks, "SELECT * FROM tasks WHERE start_at = ?", startAt); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t TaskRepository) FindByEndDate(endAt *time.Time) (*[]*Task, error) {

	tasks := &[]*Task{}

	if err := t.db.Select(tasks, "SELECT * FROM tasks WHERE end_at >= ? AND end_at IS NOT NULL", endAt); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t TaskRepository) FindStarted() (*[]*Task, error) {
	tasks := &[]*Task{}

	if err := t.db.Select(tasks, "SELECT * FROM tasks WHERE end_at IS NULL"); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t TaskRepository) FindById(id int64) (*Task, error) {
	var task Task

	if err := t.db.Get(&task, "SELECT * FROM tasks WHERE id = ? LIMIT 1", id); err != nil {
		return nil, err
	}

	return &task, nil
}

func (t TaskRepository) FindByTag(tag string) (*Task, error) {
	var task Task

	if err := t.db.Get(&task, "SELECT * FROM tasks WHERE tag = ? LIMIT 1", tag); err != nil {
		return nil, err
	}

	return &task, nil
}

func (t TaskRepository) Create(task *Task) error {

	if err := t.db.Get(
		task,
		"INSERT INTO tasks (tag, description) VALUES (?, ?) RETURNING id, start_at",
		task.GetTag(), task.GetDescription(),
	); err != nil {
		return err
	}

	return nil
}

func (t TaskRepository) UpdateEndAt(task *Task) error {
	endAt := time.Now()
	if _, err := t.db.Exec(
		"UPDATE tasks SET end_at = ? WHERE id = ?", endAt, task.Id,
	); err != nil {
		return err
	}

	task.SetEndAt(&endAt)

	return nil
}
