import React from 'react'
import ReactDOM from 'react-dom'
import CssBaseline from '@material-ui/core/CssBaseline'
import { ThemeProvider } from '@material-ui/core/styles'
import App from './views/app'
import theme from './theme'

ReactDOM.render(
  <ThemeProvider theme={theme}>
    {/* CssBaseline kickstarts an elegent baseline to build on */}
    <CssBaseline />
    <App />
  </ThemeProvider>,
  document.getElementById('root')
)
