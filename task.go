package qtodo

import (
	"errors"
	"time"
)

// TaskImpl implements the Task interface
type TaskImpl struct {
	action      func()
	alarmTime   time.Time
	name        string
	description string
}

// Task interface defines the required methods for a task
type Task interface {
	DoAction()
	GetAlarmTime() time.Time
	GetAction() func()
	GetName() string
	GetDescription() string
}

// NewTask creates a new task with the given parameters
func NewTask(action func(), alarmTime time.Time, name string, description string) (*TaskImpl, error) {
	// Validate inputs
	if action == nil {
		return nil, errors.New("action cannot be nil")
	}

	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	if description == "" {
		return nil, errors.New("description cannot be empty")
	}

	// Validate time
	if alarmTime.Before(time.Now()) {
		return nil, errors.New("alarm time cannot be in the past")
	}

	// Create and return new task
	return &TaskImpl{
		action:      action,
		alarmTime:   alarmTime,
		name:        name,
		description: description,
	}, nil
}

// DoAction executes the task's action
func (t *TaskImpl) DoAction() {
	t.action()
}

// GetAlarmTime returns the task's scheduled time
func (t *TaskImpl) GetAlarmTime() time.Time {
	return t.alarmTime
}

// GetAction returns the task's action function
func (t *TaskImpl) GetAction() func() {
	return t.action
}

// GetName returns the task's name
func (t *TaskImpl) GetName() string {
	return t.name
}

// GetDescription returns the task's description
func (t *TaskImpl) GetDescription() string {
	return t.description
}
