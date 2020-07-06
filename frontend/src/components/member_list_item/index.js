import React from 'react'
import PropTypes from 'prop-types'
import RemotePeer from '../peer_remote'

const MemberListItem = ({ id, name, localStream }) => (
  <li>
    <RemotePeer id={id} name={name} localStream={localStream} />
  </li>
)

MemberListItem.defaultProps = {
  localStream: null,
}

MemberListItem.propTypes = {
  id: PropTypes.string.isRequired,
  name: PropTypes.string.isRequired,
  localStream: PropTypes.instanceOf(MediaStream),
}

export default MemberListItem
