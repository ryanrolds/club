import React, { createContext } from 'react'
import { useDispatch } from 'react-redux'

import { socketConnected, socketDisconnected } from './store/actions/websocket'
import { ROOM_JOINED, ROOM_LEFT, roomJoined, roomLeft } from './store/actions/room'
import { GROUP_JOINED, groupJoined } from './store/actions/group'
import { PEER_JOIN, PEER_LEAVE, peerJoin, peerLeave } from './store/actions/peer'

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

    websocket.addEventListener('open', () => {
      dispatch(socketConnected())
    })

    websocket.addEventListener('close', () => {
      dispatch(socketDisconnected())
    })

    websocket.addEventListener('message', (message) => {
      let data = null
      try {
        data = JSON.parse(message.data)
      } catch (e) {
        console.error('problem parsing data', e, message.data)
        return
      }

      switch (data.type) {
        case ROOM_JOINED:
          dispatch(roomJoined(data.payload.id, data.payload.groups))
          break
        case ROOM_LEFT:
          dispatch(roomLeft())
          break
        case GROUP_JOINED:
          dispatch(groupJoined(data.payload.id, data.payload.members))
          break
        case PEER_JOIN:
          dispatch(peerJoin(data.peerId))
          break
        case PEER_LEAVE:
          dispatch(peerLeave(data.peerId))
          break
        default:
          console.error('invalid message type', data.type)
      }
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
