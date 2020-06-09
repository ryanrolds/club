import { combineReducers } from 'redux'
import websocketReducer from './websocket'
import localReducer from './local'
import roomReducer from './room'
import groupReducer from './group'
import peerReducer from './peers'

export default combineReducers({
  connected: websocketReducer,
  local: localReducer,
  room: roomReducer,
  group: groupReducer,
  peers: peerReducer,
})
