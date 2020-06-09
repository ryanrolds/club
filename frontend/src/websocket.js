import React, { createContext } from 'react'
import { useDispatch } from 'react-redux'

import { socketConnected, socketDisconnected } from './store/actions/websocket'

const WebSocketContext = createContext(null)
export { WebSocketContext }

const hostUrl = new URL(window.location)
const isHTTPS = hostUrl.protocol === 'https:'
const url = `${isHTTPS ? 'wss' : 'ws'}://${hostUrl.hostname}:3001/room`

export default ({ children }) => {
  let websocket
  let ws

  const dispatch = useDispatch()

  if (!websocket) {
    websocket = new WebSocket(url)

    websocket.addEventListener('open', (event) => {
      console.log(`connected:`, event)
      dispatch(socketConnected())
    })

    websocket.addEventListener('close', (event) => {
      console.log(`closed:`, event)
      dispatch(socketDisconnected())
    })

    websocket.addEventListener('message', (message) => {
      console.log(message)
    })

    const sendJoin = (group) => {
      websocket.send(
        JSON.stringify({ type: 'join', destId: '', payload: { group } })
      )
    }

    const sendLeave = () => {
      websocket.send(
        JSON.stringify({ type: 'leave', destId: '', payload: { reason: 'exit' } })
      )
    }

    const sendOffer = (peerId, offer) => {
      websocket.send(
        JSON.stringify({ type: 'offer', destId: peerId, payload: offer })
      )
    }

    const sendAnswer = (peerId, answer) => {
      websocket.send(
        JSON.stringify({ type: 'answer', destId: peerId, payload: answer })
      )
    }

    const sendICECandidate = (peerId, candidate) => {
      websocket.send(
        JSON.stringify({ type: 'icecandidate', destId: peerId, payload: candidate })
      )
    }

    ws = {
      websocket: websocket,
      sendJoin,
      sendLeave,
      sendOffer,
      sendAnswer,
      sendICECandidate,
    }
  }

  return <WebSocketContext.Provider value={ws}>{children}</WebSocketContext.Provider>
}
