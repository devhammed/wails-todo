import React, { useState, useEffect, useCallback } from 'react'

const Backend = window.backend.Todos
const TodosContext = React.createContext()

export const useTodos = () => {
  const contextState = React.useContext(TodosContext)

  if (contextState === null) {
    throw new Error('useTodos must be used within a TodosProvider tag.')
  }

  return contextState
}

export const TodosProvider = ({ children }) => {
  const [todos, setTodos] = useState([])

  const loadTodos = useCallback(() => Backend.All().then(setTodos), [setTodos])

  const addTodo = useCallback((text) => Backend.Add(text).then(loadTodos), [
    loadTodos
  ])

  const deleteTodo = useCallback((id) => Backend.Delete(id).then(loadTodos), [
    loadTodos
  ])

  const changeTitle = useCallback(
    (id, newTitle) => Backend.ChangeTitle(id, newTitle).then(loadTodos),
    [loadTodos]
  )

  const toggleCompleted = useCallback(
    (id) => Backend.ToggleCompleted(id).then(loadTodos),
    [loadTodos]
  )

  useEffect(loadTodos, [])

  return (
    <TodosContext.Provider
      value={{ todos, addTodo, deleteTodo, changeTitle, toggleCompleted }}
    >
      {children}
    </TodosContext.Provider>
  )
}
