import React from 'react'
import PropTypes from 'prop-types'
import Button from '@material-ui/core/Button'

const Group = ({ name, onClick }) => (
  <Button onClick={onClick}>
    Join
    {name}
  </Button>
)

Group.propTypes = {
  name: PropTypes.string.isRequired,
  onClick: PropTypes.func.isRequired,
}

export default Group
