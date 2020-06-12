export const SOCKET_CONNECTED = 'socket_connected'
export const SOCKET_DISCONNECTED = 'socket_disconnected'

export const socketConnected = () => {
  return {
    type: SOCKET_CONNECTED,
  }
}

export const socketDisconnected = () => {
  return {
    type: SOCKET_DISCONNECTED,
  }
}
