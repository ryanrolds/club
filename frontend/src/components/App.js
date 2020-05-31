import React from 'react'
import CssBaseline from '@material-ui/core/CssBaseline'
import ClubGroup from './ClubGroup.jsx'
import TopBar from './TopBar'

//   let isHTTPS = window.location.protocol !== 'https:'
// const wsServer = (isHTTPS ? "ws" : "wss") + "://localhost:3001/room"
//   return (
//     <ClubGroup URL={wsServer} />
//   )

export default function App() {
  return (
    <>
      <CssBaseline />
      <main>
        <TopBar />
        <ClubGroup URL='ws://localhost:3001/room' />
      </main>
    </>
  )
}
