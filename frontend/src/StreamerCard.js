import React, { useState, useEffect } from 'react'
import { makeStyles, Card, CardMedia } from '@material-ui/core'
import PropTypes from 'prop-types'

const StreamerCard = ({ stream, muted }) => {
  const [streamRef, setStreamRef] = useState({ streamRef: React.createRef() })
  const classes = makeStyles(() => ({
    card: {
      height: '100%',
      display: 'flex',
      flexDirection: 'column',
    },
    cardMedia: {
      paddingTop: '100%', // 16:9
    },
  }))

  useEffect(() => {
    if (streamRef.current.srcObject !== stream) {
      const newStreamRef = streamRef.current
      newStreamRef.srcObject = stream
      newStreamRef.autoplay = true
      newStreamRef.muted = muted
      setStreamRef(newStreamRef)
    }
  })

  return (
    <Card className={classes.card}>
      <CardMedia
        component='video'
        className={classes.cardMedia}
        ref={streamRef}
        title='A streamer'
      />
    </Card>
  )
}

StreamerCard.propTypes = {
  stream: PropTypes.node.isRequired,
  muted: PropTypes.bool,
}

StreamerCard.defaultProps = {
  muted: false,
}

export default StreamerCard
