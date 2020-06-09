import React, { useContext } from 'react'

import CssBaseline from '@material-ui/core/CssBaseline'
import Button from '@material-ui/core/Button';

import TopBar from '../../components/appBar/topBar'
import { WebSocketContext } from '../../websocket'

export default function App() {
  const ws = useContext(WebSocketContext);

  return (
    <>
      <CssBaseline />
      <main>
        <TopBar />
        {/* Paper or Main Page Component Here */}
        <Button variant="contained" onClick={() => { ws.sendJoin() }}>Join Default Group</Button>
      </main>
    </>
  )
}
