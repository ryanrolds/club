export const JOIN_GROUP = 'join_group'
export const LEAVE_GROUP = 'leave_group'
export const GROUP_MEMBER_JOIN = 'socket_connected'
export const GROUP_MEMBER_LEFT = 'socket_disconnected'

export const joinGroup = (id) => {
  return {
    type: JOIN_GROUP,
    id: id,
  }
}

export const leaveGroup = (id) => {
  return {
    type: LEAVE_GROUP,
    id: id,
  }
}

export const groupMemberJoined = (id) => {
  return {
    type: GROUP_MEMBER_JOIN,
    id: id,
  }
}

export const groupMemberLeft = (id) => {
  return {
    type: GROUP_MEMBER_LEFT,
    id: id,
  }
}
