import React from 'react'
import PropTypes from 'prop-types'
import RemotePeer from '../peer_remote'

const MemberListItem = ({ id, name, sendOffer, localStream }) => (
  <li>
    <RemotePeer id={id} name={name} sendOffer={sendOffer} localStream={localStream} />
  </li>
)

MemberListItem.propTypes = {
  id: PropTypes.string.isRequired,
  name: PropTypes.string.isRequired,
  sendOffer: PropTypes.bool,
  localStream: PropTypes.instanceOf(MediaStream),
}

export default MemberListItem
