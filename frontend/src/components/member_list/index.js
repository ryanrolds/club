import React, { useState } from 'react'
import PropTypes from 'prop-types'
import MemberListItem from '../member_list_item'
import LocalPeer from '../peer_local'

const MemberList = ({ localID, members }) => {
  const [localStream, setLocalStream] = useState(null)

  return (
    <ul>
      <li>
        <LocalPeer id={localID} stream={localStream} setStream={setLocalStream} />
      </li>
      {members.map(
        (member) =>
          localID !== member.id && (
            <MemberListItem
              key={member.id}
              id={member.id}
              name={member.name}
              localStream={localStream}
              sendOffer={member.sendOffer}
            />
          )
      )}
    </ul>
  )
}

MemberList.propTypes = {
  localID: PropTypes.string.isRequired,
  members: PropTypes.arrayOf(
    PropTypes.shape({
      id: PropTypes.string.isRequired,
      name: PropTypes.string.isRequired,
      sendOffer: PropTypes.bool,
    }).isRequired
  ).isRequired,
}

export default MemberList
