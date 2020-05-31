import React from 'react'
import { Paper, makeStyles } from '@material-ui/core'
import PropTypes from 'prop-types'
import ClubVideoActions from './ClubVideoActions'
import ClubVideoPlayer from './ClubVideoPlayer'

const useStyles = makeStyles(() => ({
  root: {
    transform: 'translateZ(0px)',
    flexGrow: 1,
  },
}))

function ClubMember({ mediaStream, muted }) {
  const classes = useStyles()
  return mediaStream ? (
    <Paper variant='outlined' className={classes.root}>
      <ClubVideoPlayer mediaStream={mediaStream} muted={muted} />
      <ClubVideoActions />
    </Paper>
  ) : null
}

ClubMember.defaultProps = {
  muted: false,
}

ClubMember.propTypes = {
  mediaStream: PropTypes.objectOf(PropTypes.object).isRequired,
  muted: PropTypes.bool,
}

export default ClubMember
