import { ROOM_JOINED } from '../actions/room'

const initialState = {
  id: null,
  video: false,
  audio: false,
  screenShare: false,
}

export default (state = initialState, action) => {
  const newState = { ...state }

  switch (action.type) {
    case ROOM_JOINED:
      newState.id = action.localID

      return newState
    default:
      return newState
  }
}
