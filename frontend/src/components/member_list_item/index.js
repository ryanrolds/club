import React from 'react'
import PropTypes from 'prop-types'
import RemotePeer from '../peer_remote'

const MemberListItem = ({ id }) => (
  <li>
    <RemotePeer id={id} />
  </li>
)

MemberListItem.propTypes = {
  id: PropTypes.string.isRequired,
}

export default MemberListItem
