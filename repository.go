package qtodo

import (
	"fmt"
)

type Database interface {
	GetTaskList() []Task
	GetTask(name string) (Task, error)
	SaveTask(task Task) error
	DelTask(name string) error
}

type InMemoryDatabase struct {
	tasks []Task
}

func NewInMemoryDatabase() *InMemoryDatabase {
	return &InMemoryDatabase{
		tasks: make([]Task, 0),
	}
}

func (db *InMemoryDatabase) GetTaskList() []Task {
	return db.tasks
}

func (db *InMemoryDatabase) GetTask(name string) (Task, error) {
	for _, task := range db.tasks {
		if task.GetName() == name {
			return task, nil
		}
	}
	return nil, fmt.Errorf("task not found")
}

func (db *InMemoryDatabase) SaveTask(task Task) error {
	for i, t := range db.tasks {
		if t.GetName() == task.GetName() {
			db.tasks[i] = task // Update existing task
			return nil
		}
	}
	db.tasks = append(db.tasks, task) // Add new task
	return nil
}

func (db *InMemoryDatabase) DelTask(name string) error {
	for i, task := range db.tasks {
		if task.GetName() == name {
			db.tasks = append(db.tasks[:i], db.tasks[i+1:]...) // Remove task
			return nil
		}
	}
	return fmt.Errorf("task not found")
}

func NewDatabase() Database {
	return NewInMemoryDatabase()
}
