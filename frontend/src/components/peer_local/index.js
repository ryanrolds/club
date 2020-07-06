import React, { useState } from 'react'
import PropTypes from 'prop-types'
import Button from '@material-ui/core/Button'
import { makeStyles } from '@material-ui/core/styles'

import Media from '../media'

const useStyles = makeStyles({
  root: {},
  media: {},
  button: {},
})

const PeerLocal = ({ id }) => {
  const classes = useStyles()
  const [stream, setStream] = useState(null)

  const enableVideo = () => {
    const opts = { audio: true, video: true }
    navigator.mediaDevices
      .getUserMedia(opts)
      .then((s) => {
        setStream(s)
      })
      .catch((err) => {
        console.log(err)
      })
  }

  return (
    <div className={classes.root}>
      <Media id={id} srcObject={stream} autoPlay muted className={classes.media} />
      {!stream && (
        <Button onClick={enableVideo} className={classes.button}>
          Enable video
        </Button>
      )}
    </div>
  )
}

PeerLocal.propTypes = {
  id: PropTypes.string.isRequired,
}

export default PeerLocal
