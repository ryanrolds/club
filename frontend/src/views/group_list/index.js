import React from 'react'
import PropTypes from 'prop-types'
import  GroupItem from '../group_item'

const GroupList = ({ groups }) => (
  <ul>
    {groups.map((group) => (
      <GroupItem key={group.id} name={group.name} />
    ))}
  </ul>
)

GroupList.propTypes = {
  groups: PropTypes.arrayOf(
    PropTypes.shape({
      id: PropTypes.string.isRequired,
      name: PropTypes.string.isRequire,
      // members: PropTypes.arrayOf(PropTypes.shape({}).isRequired),
      num_members: PropTypes.number.isRequired,
    }).isRequired
  ).isRequired,
}

export default GroupList
