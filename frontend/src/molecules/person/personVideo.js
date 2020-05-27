import React from 'react'
import { makeStyles } from '@material-ui/core'

const useStyles = makeStyles(() => ({
  video: {
    width: '100%',
  },
}))

function PersonVideo({ person }) {
  const classes = useStyles()

  return (
    <video
      ref={(video) => {
        video.srcObject = person.stream
      }}
      autoPlay
      muted={person.muted}
      className={classes.video}
    />
  )
}

export default PersonVideo
