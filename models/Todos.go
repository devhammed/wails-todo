package models

import (
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

// WailsInit initializes the wails runtime.
func (t *Todos) WailsInit(runtime *wails.Runtime) error {
	homedir, err := runtime.FileSystem.HomeDir()

	if err != nil {
		return err
	}

	t.db = path.Join(homedir, ".wails.todos.json")

	t.ensureDatabaseExists()

	return nil
}

// Add appends a new todo.
func (t *Todos) Add(text string) {
	t.todos = append(t.todos, Todo{ID: len(t.todos), Title: text, Completed: false})
}
