import React from 'react'
import PropTypes from 'prop-types'
import Button from '@material-ui/core/Button'
import { makeStyles } from '@material-ui/core/styles'

import Media from '../media'

const useStyles = makeStyles({
  root: {
    position: 'relative',
    overflow: 'hidden',
  },
  media: {
    'text-align': 'center',
    height: '100%',
    width: '100%',
  },
  label: {
    position: 'absolute',
    top: '0.5em',
    left: '0.5em',
  },
  button: {},
})

const PeerLocal = ({ id, stream, setStream }) => {
  const classes = useStyles()

  const enableVideo = async () => {
    const opts = { audio: true, video: true }
    const mediaStream = await navigator.mediaDevices.getUserMedia(opts)
    setStream(mediaStream)
  }

  return (
    <div className={classes.root}>
      <Media id={id} srcObject={stream} autoPlay muted className={classes.media} />
      <span className={classes.label}>
        Local&nbsp;-&nbsp;
        {id}
      </span>
      {!stream && (
        <Button onClick={enableVideo} className={classes.button}>
          Enable video
        </Button>
      )}
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
