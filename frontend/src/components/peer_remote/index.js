import React, { useState, useContext, useEffect } from 'react'
import PropTypes from 'prop-types'
import { makeStyles } from '@material-ui/core/styles'

import { WebSocketContext } from '../../websocket'
import Media from '../media'

const useStyles = makeStyles({
  root: {
    position: 'relative',
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

const config = {
  iceServers: [
    {
      urls: 'stun:stun1.l.google.com:19302',
    },
  ],
}

const PeerRemote = ({ id, localStream }) => {
  const ws = useContext(WebSocketContext)
  const classes = useStyles()
  const [stream, setStream] = useState(null)
  const [peer] = useState(new RTCPeerConnection(config))
  const tracks = []

  const getOffer = async () => {
    const offer = await peer.createOffer({
      offerToReceiveVideo: 1,
      offerToReceiveAudio: 1,
    })
    await peer.setLocalDescription(offer)
    return offer
  }

  const getAnswer = async () => {
    const answer = await peer.createAnswer()
    await peer.setLocalDescription(answer)
    return answer
  }

  const addTracks = () => {
    localStream.getTracks().forEach((track) => {
      if (tracks.indexOf(track) === -1) {
        tracks.push(track)
        peer.addTrack(track, localStream)
      }
    })
  }

  useEffect(() => {
    peer.addEventListener('icecandidate', ({ candidate }) => {
      if (candidate) {
        ws.sendICECandidate(id, candidate)
      }
    })

    peer.addEventListener('track', (track) => {
      setStream(track.streams[0])
    })

    peer.addEventListener('negotiationneeded', async () => {
      const offer = await getOffer()
      ws.sendOffer(id, offer)
    })

    const wsEventListener = async (type, event) => {
      let answer = null

      switch (type) {
        case 'offer':
          await peer.setRemoteDescription(event)

          answer = await getAnswer()
          ws.sendAnswer(id, answer)
          break
        case 'answer':
          await peer.setRemoteDescription(event)
          break
        case 'icecandidate':
          await peer.addIceCandidate(event)
          break
        default:
          console.log('unknown event', event)
      }
    }

    ws.addPeerEventListener(id, wsEventListener)

    return () => {
      ws.removePeerEventListener(id, wsEventListener)
    }
  }, []) // eslint-disable-line react-hooks/exhaustive-deps

  useEffect(() => {
    if (!localStream) {
      return
    }

    addTracks()
  }, [localStream]) // eslint-disable-line react-hooks/exhaustive-deps

  return (
    <div className={classes.root}>
      <Media id={id} srcObject={stream} autoPlay className={classes.media} />
      <span className={classes.label}>
        Peer&nbsp;-&nbps;
        {id}
      </span>
    </div>
  )
}

PeerRemote.defaultProps = {
  localStream: null,
}

PeerRemote.propTypes = {
  id: PropTypes.string.isRequired,
  localStream: PropTypes.instanceOf(MediaStream),
}

export default PeerRemote
