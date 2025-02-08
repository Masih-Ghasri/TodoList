# QTodo

QTodo is a task management library written in Go that provides functionality for creating, managing, and scheduling tasks.

## Features

- Create time-scheduled tasks
- Define custom actions for each task
- Manage active and inactive tasks
- In-memory task storage
- Optional automatic task deletion after execution

## Usage

To use this library, first create instances of `Database` and `Application`:

```go
db := qtodo.NewDatabase()
app := qtodo.NewApp(db)
```

### Adding a New Task

```go
// Create a task that will execute in 5 minutes
futureTime := time.Now().Add(5 * time.Minute)
err := app.AddTask(
    "my-task",                    // task name
    "This is a test task",        // description
    futureTime,                   // execution time
    func() {                      // action
        fmt.Println("Task executed!")
    },
    true,                         // delete after execution
)
```

### Starting a Task

```go
err := app.StartTask("my-task")
if err != nil {
    log.Fatal(err)
}
```

### Stopping a Task

```go
app.StopTask("my-task")
```

### Getting Task Lists

```go
// Get all tasks
allTasks := app.GetTaskList()

// Get active tasks
activeTasks := app.GetActiveTaskList()
```

## Project Structure

- `app.go`: Main application implementation and task management
- `repository.go`: Task storage management
- `task.go`: Task structure definition and related methods

## Implementation Notes

- Tasks cannot have empty names or descriptions
- Task execution time must be in the future
- Task names must be unique
- Tasks can be automatically deleted after execution

## Complete Example

```go
package main

import (
    "fmt"
    "time"
    "github.com/your-username/qtodo"
)

func main() {
    // Create new instances
    db := qtodo.NewDatabase()
    app := qtodo.NewApp(db)

    // Add new task
    err := app.AddTask(
        "reminder",
        "Meeting reminder",
        time.Now().Add(1 * time.Hour),
        func() {
            fmt.Println("It's time for the meeting!")
        },
        true,
    )
    
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }

    // Start the task
    err = app.StartTask("reminder")
    if err != nil {
        fmt.Printf("Error starting task: %v\n", err)
        return
    }

    // Keep the program running
    time.Sleep(2 * time.Hour)
}
```

## Requirements

- Go 1.11 or higher

## License

This project is released under the MIT License.