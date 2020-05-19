import React from 'react'
import { makeStyles, Card, CardMedia } from '@material-ui/core'

export default class StreamerCard extends React.Component {
  static useStyles() {
    return makeStyles(() => ({
      card: {
        height: '100%',
        display: 'flex',
        flexDirection: 'column',
      },
      cardMedia: {
        paddingTop: '100%', // 16:9
      },
    }))
  }

  constructor(props) {
    super(props)
    this.streamRef = React.createRef()
  }

  componentDidMount() {
    this.updateVideoStream()
  }

  componentDidUpdate() {
    this.updateVideoStream()
  }

  updateVideoStream() {
    const { stream, muted } = this.props
    if (this.streamRef.current.srcObject !== stream) {
      this.streamRef.current.srcObject = stream
      this.streamRef.current.autoplay = true
      this.streamRef.current.muted = muted
    }
  }

  render() {
    const classes = StreamerCard.useStyles()

    return (
      <Card className={classes.card}>
        <CardMedia
          component='video'
          className={classes.cardMedia}
          ref={this.streamRef}
          title='A streamer'
        />
      </Card>
    )
  }
}
