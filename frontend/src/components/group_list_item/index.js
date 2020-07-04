import React from 'react'
import PropTypes from 'prop-types'

const Group = ({ name, onClick }) => <li onClick={onClick}>{name}</li>

Group.propTypes = {
  name: PropTypes.string.isRequired,
  onClick: PropTypes.func.isRequired,
}

export default Group
