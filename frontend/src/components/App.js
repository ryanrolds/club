import React from 'react';
import CssBaseline from '@material-ui/core/CssBaseline'
import RTCMesh from './RTCMesh.jsx';
import TopBar from './TopBar'

//   let isHTTPS = window.location.protocol !== 'https:'
// const wsServer = (isHTTPS ? "ws" : "wss") + "://localhost:3001/room"
//   return (
//     <RTCMesh URL={wsServer} />
//   )

export default function App (){
  return (
    <>
      <CssBaseline />
      <main>
        <TopBar />
        <RTCMesh URL="ws://localhost:3001/room" />
      </main>
    </>
  )
}
