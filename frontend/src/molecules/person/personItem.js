import React from 'react'
import { Paper, makeStyles } from '@material-ui/core'
import PropTypes from 'prop-types'
import PersonActions from './personActions'
import PersonVideo from './personVideo'

const useStyles = makeStyles((theme) => ({
  root: {
    transform: 'translateZ(0px)',
    flexGrow: 1,
  },
}))

function PersonItem({ person }) {
  const classes = useStyles()
  const peerConnection = new RTCPeerConnection()

  peerConnection.addEventListener('icecandidate', ({candidate}) => {
    if (candidate) {
      // add ice candidate to peers
    }
  })

  // person.stream.forEach((track) => {
  //   peerConnection.addTrack(track, person.stream)
  // })

  // const offer = await peerConnection.createOffer({
  //   offerToReceiveVideo: 1,
  //   offerToReceiveAudio: 1,
  // })

  // await peerConnection.setLocalDescription(offer)

  return (
    <Paper
      variant="outlined"
      className={classes.root}
    >
      <PersonVideo person={person}/>
      <PersonActions />
    </Paper>
  )
}

PersonItem.propTypes = {
  person: PropTypes.object.isRequired
}

export default PersonItem
