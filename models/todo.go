package models

import "time"

type Todo struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"is_completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Database
var Todos []Todo

// GetAllTodos returns all todos
func GetAllTodos() []Todo {
	return Todos
}

// GetTodoByID finds a todo by ID, returns it and its index
func GetTodoByID(id int) (Todo, int) {
	for i, todo := range Todos {
		if todo.ID == id {
			return todo, i
		}
	}
	return Todo{}, -1
}

// CreateTodo adds a new todo to the slice
func CreateTodo(id int, title string, description string) Todo {
	todo := Todo{
		ID:          id,
		Title:       title,
		Description: description,
		IsCompleted: false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	Todos = append(Todos, todo)
	return todo
}

// UpdateTodo replaces a todo fully
func UpdateTodo(index int, title, description string, isCompleted bool) Todo {
	Todos[index].Title = title
	Todos[index].Description = description
	Todos[index].IsCompleted = isCompleted
	Todos[index].UpdatedAt = time.Now()
	return Todos[index]
}

// PatchTodo updates only the fields that are provided
func PatchTodo(index int, fields map[string]interface{}) Todo {
	if val, ok := fields["title"]; ok {
		Todos[index].Title = val.(string)
	}
	if val, ok := fields["description"]; ok {
		Todos[index].Description = val.(string)
	}
	if val, ok := fields["is_completed"]; ok {
		Todos[index].IsCompleted = val.(bool)
	}
	Todos[index].UpdatedAt = time.Now()
	return Todos[index]
}

// DeleteTodo
func DeleteTodo(id int) {
    // Create a new empty slice
    var newTodos []Todo

    for _, todo := range Todos {
        if todo.ID == id {
            continue
        }
        newTodos = append(newTodos, todo)
    }

    // Replace the original slice with the filtered slice
    Todos = newTodos
}