import React from 'react'
import { makeStyles } from '@material-ui/core'
import PropTypes from 'prop-types'

const useStyles = makeStyles(() => ({
  video: {
    width: '100%',
    backgroundColor: 'grey'
  }
}))

function RTCVideoPlayer({mediaStream, muted}) {
  const classes = useStyles()
  const videoRef = React.createRef()

  const setRef = () => {
    if(!mediaStream) return
    videoRef.srcObject = mediaStream
  }

  return (
    <video
      className={classes.video}
      autoPlay
      muted={muted}
      ref={setRef()}
    >
      <track default />
    </video>
  )
}

RTCVideoPlayer.defaultProps = {
  muted: false
}

RTCVideoPlayer.propTypes = {
  mediaStream: PropTypes.objectOf(PropTypes.object).isRequired,
  muted: PropTypes.bool
}

export default RTCVideoPlayer
