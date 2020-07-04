import React from 'react'
import PropTypes from 'prop-types'

const Group = ({ name }) => (
  <li>
    {name}
  </li>
)

Group.propTypes = {
  name: PropTypes.string.isRequired
}

export default Group