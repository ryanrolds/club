import React, { createContext } from 'react'
import { useDispatch } from 'react-redux'

import lodash from 'lodash'

import { socketConnected, socketDisconnected } from './store/actions/websocket'
import { ROOM_JOINED, ROOM_LEFT, roomJoined, roomLeft } from './store/actions/room'
import { GROUP_JOINED, groupJoined } from './store/actions/group'
import { PEER_JOIN, PEER_LEAVE, peerJoin, peerLeave } from './store/actions/peer'
import { RTC_OFFER, RTC_ANSWER, RTC_ICECANDIDATE } from './store/actions/rtc'

const WebSocketContext = createContext(null)
export { WebSocketContext }

const hostUrl = new URL(window.location)
const isHTTPS = hostUrl.protocol === 'https:'
const url = `${isHTTPS ? 'wss' : 'ws'}://${hostUrl.hostname}:3001/room`

export default ({ children }) => {
  let websocket
  let ws
  const peerEventListeners = {}
  const dispatch = useDispatch()

  if (!websocket) {
    const addPeerEventListener = (peerID, listener) => {
      if (!lodash.has(peerEventListeners, peerID)) {
        peerEventListeners[peerID] = []
      }

      peerEventListeners[peerID].push(listener)
    }

    const removePeerEventListener = (peerID, listener) => {
      if (!lodash.has(peerEventListeners, peerID)) {
        throw new Error('unknown peerID')
      }

      peerEventListeners[peerID] = lodash.remove(
        peerEventListeners[peerID],
        listener
      )
    }

    const notifyPeerEventListeners = (peerID, type, event) => {
      if (!lodash.has(peerEventListeners, peerID)) {
        throw new Error('unknown peerID')
      }

      peerEventListeners[peerID].forEach((listener) => {
        listener(type, event)
      })
    }

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
          dispatch(
            roomJoined(data.payload.id, data.payload.local_id, data.payload.groups)
          )
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
        case RTC_OFFER:
        case RTC_ANSWER:
        case RTC_ICECANDIDATE:
          notifyPeerEventListeners(data.peerId, data.type, data.payload)
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
      addPeerEventListener,
      removePeerEventListener,
    }
  }

  return <WebSocketContext.Provider value={ws}>{children}</WebSocketContext.Provider>
}
