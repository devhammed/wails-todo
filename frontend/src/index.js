import 'core-js/stable'
import React from 'react'
import ReactDOM from 'react-dom'
import * as Wails from '@wailsapp/runtime'

import './index.css'
import App from './containers/App'

Wails.Init(() => {
  ReactDOM.render(<App />, document.getElementById('app'))
})
