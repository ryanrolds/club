import { GROUP_JOINED } from '../actions/group'

const intialState = {}

export default (state = intialState, action) => {
  switch (action.type) {
    case GROUP_JOINED:
      return {
        id: action.id,
        members: action.members,
      }
    default:
      return state
  }
}
