import websocketReducer, { CONNECTED, DISCONNECTED } from './websocket'
import { socketConnected, socketDisconnected } from '../actions/websocket'

it('should return true when connected', () => {
  const state = websocketReducer({}, socketConnected())
  expect(state).toEqual(CONNECTED)
})

it('should return false when disconnected', () => {
  const state = websocketReducer({}, socketDisconnected())
  expect(state).toEqual(DISCONNECTED)
})
