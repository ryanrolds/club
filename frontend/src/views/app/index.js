import React, { useContext } from 'react'

import CssBaseline from '@material-ui/core/CssBaseline'
import Button from '@material-ui/core/Button'

import TopBar from '../../components/appBar/topBar'
import { WebSocketContext } from '../../websocket'

import store from '../../store'

let currentState = store.getState()

export default function App() {
  const ws = useContext(WebSocketContext)

  store.subscribe(() => {
    currentState = store.getState()
    // TODO do something with state change
    console.log(currentState)
  })

  return (
    <>
      <CssBaseline />
      <main>
        <TopBar />
        {/* Paper or Main Page Component Here */}
        <Button
          variant='contained'
          onClick={() => {
            ws.sendJoin()
          }}
        >
          Join Default Group
        </Button>
      </main>
    </>
  )
}
