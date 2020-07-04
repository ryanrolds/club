import { ROOM_JOINED, ROOM_LEFT } from '../actions/room'

const defaultState = {
  id: null,
  groups: [],
}

export default (state = defaultState, action) => {
  switch (action.type) {
    case ROOM_JOINED:
      return {
        id: action.id,
        groups: action.groups,
      }
    case ROOM_LEFT:
      return {}
    default:
      return state
  }
}
