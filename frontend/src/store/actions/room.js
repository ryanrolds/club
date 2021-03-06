export const ROOM_JOINED = 'room_joined'
export const ROOM_LEFT = 'room_left'

export const roomJoined = (id, localID, groups) => {
  return {
    type: ROOM_JOINED,
    id,
    localID,
    groups,
  }
}

export const roomLeft = () => {
  return {
    type: ROOM_LEFT,
  }
}
