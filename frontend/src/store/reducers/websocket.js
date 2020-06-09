import { SOCKET_CONNECTED, SOCKET_DISCONNECTED } from '../actions/websocket'

export default (state = false, action) => {
  switch (action.type) {
    case SOCKET_CONNECTED:
      return true
    case SOCKET_DISCONNECTED:
      return false
    default:
      return state
  }
}
