import roomReducer from './room'
import { roomJoined, roomLeft } from '../actions/room'

const groupFoo = {
  id: 'foo',
  limit: 12,
  name: 'foo',
  num_members: 0,
}

it('should return room joined state', () => {
  const state = roomReducer({}, roomJoined('id', [groupFoo]))
  expect(state).toEqual({
    id: 'id',
    groups: [groupFoo],
  })
})

it('should return room left state', () => {
  const currentState = {
    id: 'id',
    groups: [groupFoo],
  }

  const state = roomReducer(currentState, roomLeft())
  expect(state).toEqual({})
})
