import React, { createRef } from 'react'
import { makeStyles } from '@material-ui/core'

const useStyles = makeStyles((theme) => ({
  root: {
    paddingTop: '100%',
  }
}))

function PersonVideo({srcObject}){
  const classes = useStyles()
  const streamRef = createRef()
  if(streamRef.current && streamRef.current.srcObject !== srcObject) {
    streamRef.srcObject = srcObject
    streamRef.autoplay = true
    streamRef.muted = true
  }

  return (
    <video
      className={classes.root}
      ref={streamRef}
      />
  )
}

export default PersonVideo
