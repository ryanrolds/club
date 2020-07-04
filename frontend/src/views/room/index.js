import React, { useContext } from 'react'
import { connect } from 'react-redux'
import PropTypes from 'prop-types'
import GroupList from '../../components/group_list'
import { WebSocketContext } from '../../websocket'

const Room = ({ groups }) => {
  const ws = useContext(WebSocketContext)

  return <GroupList groups={groups} onGroupClick={(id) => ws.sendJoin(id)} />
}

Room.propTypes = {
  groups: PropTypes.arrayOf(
    PropTypes.shape({
      id: PropTypes.string.isRequired,
      name: PropTypes.string.isRequire,
      // members: PropTypes.arrayOf(PropTypes.shape({}).isRequired),
      num_members: PropTypes.number.isRequired,
    }).isRequired
  ).isRequired,
}

const mapStateToProps = (state) => {
  return {
    groups: state.room.groups,
  }
}

export default connect(mapStateToProps)(Room)
