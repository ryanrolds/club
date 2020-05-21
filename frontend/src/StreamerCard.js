import React from 'react'
import { makeStyles, Card, CardMedia } from '@material-ui/core'
import PropTypes from 'prop-types'

const StreamerCard = ({ stream, muted }) => {
  const streamRef = React.createRef()
  if (streamRef.current && streamRef.current.srcObject !== stream) {
    streamRef.srcObject = stream
    streamRef.autoplay = true
    streamRef.muted = muted
  }

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
  stream: PropTypes.func.isRequired,
  muted: PropTypes.bool,
}

StreamerCard.defaultProps = {
  muted: false,
}

export default StreamerCard
