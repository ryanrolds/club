import { GROUP_JOINED } from '../actions/group'
import { PEER_JOIN, PEER_LEAVE } from '../actions/peer'

const intialState = {}

export default (state = intialState, action) => {
  const newState = { ...state }

  switch (action.type) {
    case GROUP_JOINED:
      return {
        id: action.id,
        members: action.members,
      }
    case PEER_JOIN:
      // Peers joining and leaving a room can also trigger this event,
      // ignore the event if we are not in a group
      if (state.id) {
        newState.members = [
          ...state.members.slice(0),
          { id: action.id, name: action.id },
        ]
      }

      return newState
    case PEER_LEAVE:
      // Peers joining and leaving a room can also trigger this event,
      // ignore the event if we are not in a group
      if (state.id) {
        newState.members = state.members.filter((member) => {
          return member.id !== action.id
        })
      }

      return newState
    default:
      return state
  }
}
