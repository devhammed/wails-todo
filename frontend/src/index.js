import 'core-js/stable'
import React from 'react'
import ReactDOM from 'react-dom'
import * as Wails from '@wailsapp/runtime'

import App from './components/App'

Wails.Init(() => {
  ReactDOM.render(<App />, document.getElementById('app'))
})
