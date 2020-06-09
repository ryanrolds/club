import React from 'react'
import ReactDOM from 'react-dom'

import { Provider } from 'react-redux'

import CssBaseline from '@material-ui/core/CssBaseline'
import { ThemeProvider } from '@material-ui/core/styles'

import store from './store'

import App from './views/app'
import theme from './theme'

ReactDOM.render(
  <Provider store={store}>
    <WebsocketProvider>
      <ThemeProvider theme={theme}>
        {/* CssBaseline kickstarts an elegent baseline to build on */}
        <CssBaseline />
        <App />
      </ThemeProvider>
    </WebsocketProvider>
  </Provider>,
  document.getElementById('root')
)
