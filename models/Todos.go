package models

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	"github.com/wailsapp/wails"
)

// Todos is the list of what needs to be done.
type Todos struct {
	todos []Todo
	db    string
}

// NewTodos attempts to create a new Todo list.
func NewTodos() *Todos {
	return &Todos{}
}

// ensureDatabaseExists makes sure the flat db is ready.
func (t *Todos) ensureDatabaseExists() {
	_, err := os.Stat(t.db)

	if os.IsNotExist(err) {
		ioutil.WriteFile(t.db, []byte("[]"), 0600)
	}
}

// loadTodos populates the todos from the flat-file database.
func (t *Todos) loadTodos() {
	bytes, err := ioutil.ReadFile(t.db)

	if err != nil {
		return
	}

	var data []Todo

	if err := json.Unmarshal(bytes, &data); err != nil {
		return
	}

	t.todos = data
}

// saveTodos saves the todos to the flat-file database.
func (t *Todos) saveTodos() {
	todos, err := json.Marshal(t.todos)

	if err != nil {
		return
	}

	ioutil.WriteFile(t.db, []byte(todos), 0600)
}

// WailsInit initializes the wails runtime.
func (t *Todos) WailsInit(runtime *wails.Runtime) error {
	homedir, err := runtime.FileSystem.HomeDir()

	if err != nil {
		return err
	}

	t.db = path.Join(homedir, ".wails.todos.json")
	t.ensureDatabaseExists()
	t.loadTodos()

	return nil
}

// All returns all the todos.
func (t *Todos) All() []Todo {
	return t.todos
}

// Add appends a new todo.
func (t *Todos) Add(text string) {
	t.todos = append(t.todos, Todo{ID: len(t.todos), Title: text, Completed: false})
	t.saveTodos()
}

// Delete deletes a todo.
func (t *Todos) Delete(id string) []Todo {
	return t.todos
}

// ChangeTitle updates a todo's title.
func (t *Todos) ChangeTitle(id string, title string) []Todo {
	return t.todos
}

// ToggleCompleted toggles a todo's completed.
func (t *Todos) ToggleCompleted(id string) []Todo {
	return t.todos
}
