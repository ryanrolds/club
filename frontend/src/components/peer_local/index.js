import React, { useState } from 'react'
import PropTypes from 'prop-types'
import { ButtonGroup, Button } from '@material-ui/core'
import { makeStyles } from '@material-ui/core/styles'

import Media from '../media'

const ENABLE = true
const DISABLE = false

const useStyles = makeStyles({
  root: {
    position: 'relative',
    overflow: 'hidden',
    'background-color': '#000000',
    height: '100%',
  },
  label: {
    position: 'absolute',
    top: '0.5em',
    left: '0.5em',
  },
  media: {
    'text-align': 'center',
    height: '100%',
    width: '100%',
  },
  controls: {
    position: 'absolute',
    bottom: '0.5em',
    left: '0.5em',
  },
  videoButton: {},
  audioButton: {},
})

const PeerLocal = ({ id, stream, setStream }) => {
  const classes = useStyles()
  const [video, setVideo] = useState(false)
  const [audio, setAudio] = useState(false)

  const toggleTracks = (tracks, status) => {
    tracks.forEach((track) => {
      const t = track
      t.enabled = status
    })
  }

  const getStream = async () => {
    if (stream) {
      return stream
    }

    const opts = { audio: true, video: true }
    const mediaStream = await navigator.mediaDevices.getUserMedia(opts)

    // Start all tracks disabled
    toggleTracks(mediaStream.getTracks(), DISABLE)
    setStream(mediaStream)

    return mediaStream
  }

  const enableVideo = async () => {
    const mediaStream = await getStream()
    toggleTracks(mediaStream.getVideoTracks(), ENABLE)
    setVideo(ENABLE)
  }

  const disableVideo = async () => {
    const mediaStream = await getStream()
    toggleTracks(mediaStream.getVideoTracks(), DISABLE)
    setVideo(DISABLE)
  }

  const enableAudio = async () => {
    const mediaStream = await getStream()
    toggleTracks(mediaStream.getAudioTracks(), ENABLE)
    setAudio(ENABLE)
  }

  const disableAudio = async () => {
    const mediaStream = await getStream()
    toggleTracks(mediaStream.getAudioTracks(), DISABLE)
    setAudio(DISABLE)
  }

  return (
    <div className={classes.root}>
      <Media id={id} srcObject={stream} autoPlay muted className={classes.media} />
      <span className={classes.label} color='primary'>
        Local&nbsp;-&nbsp;
        {id}
      </span>
      <ButtonGroup
        size='small'
        aria-label='small outlined button group'
        color='primary'
        className={classes.controls}
      >
        {!video && (
          <Button onClick={enableVideo} className={classes.videoButton}>
            Enable video
          </Button>
        )}
        {video && (
          <Button onClick={disableVideo} className={classes.videoButton}>
            Disable video
          </Button>
        )}
        {!audio && (
          <Button onClick={enableAudio} className={classes.audioButton}>
            Enable Audio
          </Button>
        )}
        {audio && (
          <Button onClick={disableAudio} className={classes.audioButton}>
            Disable Audio
          </Button>
        )}
      </ButtonGroup>
    </div>
  )
}

PeerLocal.defaultProps = {
  stream: null,
}

PeerLocal.propTypes = {
  id: PropTypes.string.isRequired,
  stream: PropTypes.instanceOf(MediaStream),
  setStream: PropTypes.func.isRequired,
}

export default PeerLocal
