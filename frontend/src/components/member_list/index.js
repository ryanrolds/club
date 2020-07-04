import React from 'react'
import PropTypes from 'prop-types'
import MemberListItem from '../member_list_item'

const MemberList = ({ members }) => (
  <ul>
    {members.map((member) => (
      <MemberListItem key={member.id} id={member.id} name={member.name}/>
    ))}
  </ul>
)

MemberList.propTypes = {
  members: PropTypes.arrayOf(
    PropTypes.shape({
      id: PropTypes.string.isRequired,
      name: PropTypes.string.isRequired,
    }).isRequired
  ).isRequired,
}

export default MemberList
