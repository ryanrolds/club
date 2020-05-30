import React from 'react'
import { Paper, makeStyles } from '@material-ui/core'
import PropTypes from 'prop-types'
import RTCVideoActions from './RTCVideoActions'
import RTCVideoPlayer from './RTCVideoPlayer'

const useStyles = makeStyles(() => ({
  root: {
    transform: 'translateZ(0px)',
    flexGrow: 1,
  },
}))

function RTCVideo({ mediaStream, muted }) {
  const classes = useStyles()
  return mediaStream ? (
    <Paper variant="outlined" className={classes.root}>
      <RTCVideoPlayer mediaStream={mediaStream} muted={muted} />
      <RTCVideoActions />
    </Paper>
  ) : null
}

RTCVideo.propTypes = {
  mediaStream: PropTypes.objectOf(PropTypes.object).isRequired,
  muted: PropTypes.bool
}

export default RTCVideo
