import React from 'react'
import { Paper, makeStyles } from '@material-ui/core'
import PropTypes from 'prop-types'
import PersonActions from './personActions'
import PersonVideo from './personVideo'

const useStyles = makeStyles((theme) => ({
  root: {
    transform: 'translateZ(0px)',
    flexGrow: 1,
  },
}))

function PersonItem({ person }) {
  const classes = useStyles()
  return person.stream ? (
    <Paper variant='outlined' className={classes.root}>
      <PersonVideo person={person} />
      <PersonActions />
    </Paper>
  ) : null
}

PersonItem.propTypes = {
  person: PropTypes.object.isRequired,
}

export default PersonItem
