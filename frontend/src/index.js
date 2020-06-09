import React from 'react'
import ReactDOM from 'react-dom'

import { Provider } from 'react-redux'

import CssBaseline from '@material-ui/core/CssBaseline'
import { ThemeProvider } from '@material-ui/core/styles'

import store from './store'
import WebSocketProvider from './websocket';

import App from './views/app'
import theme from './theme'


ReactDOM.render(
  <Provider store={store}>
    <WebSocketProvider>
      <ThemeProvider theme={theme}>
        {/* CssBaseline kickstarts an elegent baseline to build on */}
        <CssBaseline />
        <App />
      </ThemeProvider>
    </WebSocketProvider>
  </Provider>,
  document.getElementById('root')
)
