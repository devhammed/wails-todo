import React from 'react'

import { useTodos } from '../services/Todos'

export default function TodoList () {
  const { todos } = useTodos()

  return (
    <>
      {todos.map((todo) => (
        <p key={todo.id}>{todo.title}</p>
      ))}
    </>
  )
}
