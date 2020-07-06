import React, { useState, useContext, useEffect } from 'react'
import PropTypes from 'prop-types'
import { makeStyles } from '@material-ui/core/styles'

import { WebSocketContext } from '../../websocket'
import Media from '../media'

const useStyles = makeStyles({
  root: {},
  media: {},
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

  const getOffer = () => {
    return peer
      .createOffer({
        offerToReceiveVideo: 1,
        offerToReceiveAudio: 1,
      })
      .then((offer) => {
        return peer.setLocalDescription(offer).then(() => {
          return offer
        })
      })
  }

  const getAnswer = () => {
    return peer.createAnswer().then((answer) => {
      return peer.setLocalDescription(answer).then(() => {
        return answer
      })
    })
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

    peer.addEventListener('negotiationneeded', () => {
      getOffer()
        .then((offer) => {
          ws.sendOffer(id, offer)
        })
        .catch((err) => {
          console.log(err)
        })
    })

    ws.addPeerEventListener(id, (type, event) => {
      switch (type) {
        case 'offer':
          peer.setRemoteDescription(event)

          getAnswer()
            .then((offer) => {
              ws.sendAnswer(id, offer)
            })
            .catch((err) => {
              console.log(err)
            })
          break
        case 'answer':
          peer.setRemoteDescription(event)
          break
        case 'icecandidate':
          peer.addIceCandidate(event)
          break
        default:
          console.log('unknown event', event)
      }
    })
  }, []) // eslint-disable-line react-hooks/exhaustive-deps

  useEffect(() => {
    if (!localStream) {
      return
    }

    addTracks()
  }, [localStream]) // eslint-disable-line react-hooks/exhaustive-deps

  return (
    <div className={classes.root}>
      <Media id={id} srcObject={stream} autoPlay muted className={classes.media} />
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
