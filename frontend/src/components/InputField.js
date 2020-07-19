import React, { useState } from 'react'

import { useTodos } from '../services/Todos'

export default function InputField () {
  const [text, setText] = useState('')
  const { addTodo } = useTodos()

  const handleChange = (e) => setText(e.target.value)
  const handleSubmit = (e) => {
    e.preventDefault()
    addTodo(text)
    setText('')
  }

  return (
    <form onSubmit={handleSubmit}>
      <input onChange={handleChange} />
    </form>
  )
}
