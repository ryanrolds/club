export const PEER_JOIN = 'join'
export const PEER_LEAVE = 'leave'

export const peerJoin = (id) => {
  return {
    type: PEER_JOIN,
    id: id,
  }
}

export const peerLeave = (id) => {
  return {
    type: PEER_LEAVE,
    id: id,
  }
}
