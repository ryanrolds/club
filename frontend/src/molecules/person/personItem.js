import React from 'react'
import { Paper, makeStyles } from '@material-ui/core'
import PropTypes from 'prop-types'
import PersonActions from './personActions'
import PersonVideo from './personVideo'

const useStyles = makeStyles((theme) => ({
  root: {
    height: '100%',
    transform: 'translateZ(0px)',
    flexGrow: 1,
    display: 'flex',
    flexDirection: 'column',
  },
}))

function PersonItem({ personId }) {
  const classes = useStyles()

  return (
    <Paper
      variant="outlined"
      className={classes.root}
    >
      <PersonVideo />
      <PersonActions />
    </Paper>
  )
}

PersonItem.propTypes = {
  personId: PropTypes.any.isRequired
}

export default PersonItem
