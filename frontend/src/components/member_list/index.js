import React from 'react'
import PropTypes from 'prop-types'
import MemberListItem from '../member_list_item'
import LocalPeer from '../peer_local'

const MemberList = ({ localID, members }) => (
  <ul>
    <li>
      <LocalPeer />
    </li>
    {members.map(
      (member) =>
        localID !== member.id && (
          <MemberListItem key={member.id} id={member.id} name={member.name} />
        )
    )}
  </ul>
)

MemberList.propTypes = {
  localID: PropTypes.string.isRequired,
  members: PropTypes.arrayOf(
    PropTypes.shape({
      id: PropTypes.string.isRequired,
      name: PropTypes.string.isRequired,
    }).isRequired
  ).isRequired,
}

export default MemberList
