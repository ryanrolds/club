import React, { createContext } from 'react'
import { socketConnected, socketDisconnected } from '../store/actions/websocket'

import { useDispatch } from 'react-redux';

const WebSocketContext = createContext(null)
export { WebSocketContext }

const hostUrl = new URL(location.host);
const isHTTPS = hostUrl.protocol === 'https:'
const url = (isHTTPS ? "wss" : "ws") + "://" + hostUrl.hostname + ":3001" + "/room")

export default ({ children }) => {
  let websocket
  const dispatch = useDispatch();

  if (!websocket) {
    websocket = new WebSocket(url)

    websocket.addEventListener("open", (event) => {
      console.log(`connected:`, event)
      dispatch(socketConnected())
    })

    websocket.addEventListener("close", (event) => {
      console.log(`closed:`, event)
      dispatch(socketDisonnected())
    })

    websocket.addEventListener("message", (message) => {

    })
  }

  return (
    <WebSocketContext.Provider value={ws}>
        {children}
    </WebSocketContext.Provider>
  )
}
