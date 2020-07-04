import React from 'react'
import PropTypes from 'prop-types'

const MemberListItem = ({ id }) => <li>{id}</li>

MemberListItem.propTypes = {
  id: PropTypes.string.isRequired,
}

export default MemberListItem
