import React, { createRef } from 'react'
import { makeStyles } from '@material-ui/core'

const useStyles = makeStyles(() => ({
  video: {
    width: '100%',
  }
}))

function PersonVideo({person}){
  const classes = useStyles()
  const streamRef = createRef()
  if(streamRef.current && streamRef.current.srcObject !== person.stream) {
    streamRef.current.srcObject = person.stream
    streamRef.current.autoplay = true
    streamRef.current.muted = true
  }

  console.log(streamRef)

  return (
    <video
      ref={streamRef}
      autoPlay
      className={classes.video}
      />
  )
}

export default PersonVideo
