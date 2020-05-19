import React from 'react'
import { makeStyles, Card, CardMedia } from '@material-ui/core'

export default StreamerCard((props) => {
  this.state = {
    streamRef: React.createRef()
  }
  this.useEffect = this.useEffect.bind(this);
})

StreamerCard.prototype.render(() => {
  const { streamRef } = this.state

  if(!streamRef)
    return

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
})

StreamerCard.prototype.useEffect(() => {
  const { stream, muted } = this.props

  if (!this.state.streamRef)
    return

  const { current } = this.state.streamRef

  if (current.srcObject !== stream) {
    current.srcObject = stream
    current.autoplay = true
    current.muted = muted
  }
})

Object.setPrototypeOf(StreamerCard.prototype, React.Component.prototype);
