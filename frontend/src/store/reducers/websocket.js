import { SOCKET_CONNECTED, SOCKET_DISCONNECTED } from '../actions/websocket'

export const CONNECTED = 'connected'
export const DISCONNECTED = 'disconnected'

export default (state = DISCONNECTED, action) => {
  switch (action.type) {
    case SOCKET_CONNECTED:
      return CONNECTED
    case SOCKET_DISCONNECTED:
      return DISCONNECTED
    default:
      return state
  }
}
