import React from 'react'
import { Paper, makeStyles } from '@material-ui/core'
import PropTypes from 'prop-types'
import PersonActions from './personActions'
import PersonVideo from './personVideo'

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
  },
}))

function PersonCard({ personId }) {
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

PersonCard.propTypes = {
  personId: PropTypes.any.isRequired
}

export default PersonCard
