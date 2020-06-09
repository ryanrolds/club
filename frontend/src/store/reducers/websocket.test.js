import websocketReducer from './websocket'
import { socketConnected, socketDisonnected } from '../actions/websocket'

it('should return true when connected', () => {
  let state = websocketReducer({}, socketConnected())
  expect(state).toEqual(true)
})

it('should return false when disconnected', () => {
  let state = websocketReducer({}, socketDisonnected())
  expect(state).toEqual(false)
})
