export const GROUP_JOINED = 'group_joined'
export const GROUP_LEFT = 'group_left'
export const GROUP_MEMBER_JOIN = 'peer_joined'
export const GROUP_MEMBER_LEFT = 'peed_left'

export const groupJoined = (id, members) => {
  return {
    type: GROUP_JOINED,
    id: id,
    members: members,
  }
}

export const groupLeft = (id) => {
  return {
    type: GROUP_LEFT,
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
