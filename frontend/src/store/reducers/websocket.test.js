import websocketReducer from './websocket'
import { socketConnected, socketDisconnected } from '../actions/websocket'

it('should return true when connected', () => {
  const state = websocketReducer({}, socketConnected())
  expect(state).toEqual(true)
})

it('should return false when disconnected', () => {
  const state = websocketReducer({}, socketDisconnected())
  expect(state).toEqual(false)
})
