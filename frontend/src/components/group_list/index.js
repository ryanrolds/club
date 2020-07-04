import React from 'react'
import PropTypes from 'prop-types'
import GroupListItem from '../group_list_item'

const GroupList = ({ groups, onGroupClick }) => (
  <ul>
    {groups.map((group) => (
      <GroupListItem key={group.id} name={group.name} onClick={() => onGroupClick(group.id)} />
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
    }).isRequired,
  ).isRequired,
  onGroupClick: PropTypes.func.isRequired,
}

export default GroupList
