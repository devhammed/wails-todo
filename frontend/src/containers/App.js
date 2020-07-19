import React from 'react'

import TodoList from '../components/TodoList'
import InputField from '../components/InputField'
import { TodosProvider } from '../services/Todos'

function App () {
  return (
    <TodosProvider>
      <InputField />
      <TodoList />
    </TodosProvider>
  )
}

export default App
