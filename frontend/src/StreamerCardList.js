import React from 'react'
import { makeStyles, Container, Grid } from '@material-ui/core'
import StreamerCard from './StreamerCard'
import SignalingServer from './helpers/signaling'

export default StreamerCardList(() => {
  this.state = {
    peers: {},
    signals: new SignalingServer(),
    stream: null,
    offer: null,
  }

  this.setupSignalEventHandlers = this.setupSignalEventHandlers.bind(this)
  this.onConnected = this.onConnected.bind(this)
  this.onDisconnected = this.onDisconnected.bind(this)
  this.onJoin = this.onJoin.bind(this)
  this.onOffer = this.onOffer.bind(this)
  this.onAnswer = this.onAnswer.bind(this)
  this.onICECandidate = this.onICECandidate.bind(this)
  this.newPeer = this.newPeer.bind(this)
  this.getPeer = this.getPeer.bind(this)
})

StreamerCardList.prototype.render(() => {
  const { stream, peers }

  const classes = makeStyles((theme) => ({
    cardGrid: {
      paddingTop: theme.spacing(1),
      paddingBottom: theme.spacing(1),
    },
  }))

  return (
    <Container className={classes.cardGrid}>
      <Grid container spacing={1}>
        <Grid item key={stream} xs={12} sm={6}>
          <StreamerCard stream={stream} muted />
        </Grid>
        {peers.length &&
          peers.map((peer) => (
            <Grid item key={peer} xs={12} sm={6}>
              <StreamerCard stream={peer.srcObject} />
            </Grid>
          ))}
      </Grid>
    </Container>
  )
})

StreamerCardList.prototype.useEffect(async () => {
  const opts = { audio: true, video: true }
  const stream = await navigator.mediaDevices.getUserMedia(opts)
  this.setState({ stream })
  this.setupSignalEventHandlers()
})

StreamerCardList.prototype.setupSignalEventHandlers(async () => {
  const { signals } = this.state

  signals.addEventListener('connected', async (event) => {
    console.log('connected to signalling server', event)
    await this.onConnected()
    // await peers.onConnected()

    signals.sendJoin()
  })

  signals.addEventListener('join', async (event) => {
    console.log('got join', event.join)
    let offer = await this.onJoin(event.detail)
    signals.sendOffer(event.detail.peerId, offer)
  })

  signals.addEventListener('offer', async (event) => {
    console.log('got offer', event)
    let answer = await this.onOffer(event.detail)
    signals.sendAnswer(event.detail.peerId, answer)
  })

  signals.addEventListener('icecandidate', async (event) => {
    console.log('got ice candidate', event)
    this.onICECandidate(event.detail)
  })

  signals.addEventListener('answer', async (event) => {
    console.log('got answer', event)
    await this.onAnswer(event.detail)
  })

  signals.addEventListener('leave', async (event) => {
    console.log('got leave', event)
    await this.onLeave(event.detail)
  })

  signals.addEventListener('disconnected', async (event) => {
    console.log('disconnected from signalling server', event)

    await this.onDisconnected()
    // await this.onDisconnected()
  })

  let isHTTPS = window.location.protocol !== 'https:'
  signals.connect(
    (isHTTPS ? 'ws' : 'wss') + '://' + window.location.host + '/room'
  )
})


StreamerCardList.prototype.onConnected(async () => { })

StreamerCardList.prototype.onDisconnected(async () => { })

StreamerCardList.prototype.onJoin(async (join) => {
  const { stream } = this.state
  let peer = this.getPeer(join.peerId)

  stream.getTracks().forEach((track) => {
    peer.addTrack(track, stream)
  })

  const offer = await peer.createOffer({
    offerToReceiveVideo: 1,
    offerToReceiveAudio: 1,
  })

  await peer.setLocalDescription(offer)

  return offer
})

StreamerCardList.prototype.onOffer(async (offer) => {
  const { stream } = this.state
  let peer = this.getPeer(offer.peerId)
  peer.setRemoteDescription(offer.offer)

  stream.getTracks().forEach((track) => {
    peer.addTrack(track, stream)
  })

  const answer = await peer.createAnswer()
  peer.setLocalDescription(answer)

  return answer
})

StreamerCardList.prototype.onAnswer(async (answer) => {
  let peer = this.getPeer(answer.peerId)
  await peer.setRemoteDescription(answer.answer)
})

StreamerCardList.prototype.onICECandidate(async (candidate) => {
  let peer = this.getPeer(candidate.peerId)
  peer.addIceCandidate(candidate.candidate)
})

StreamerCardList.prototype.newPeer((peerId) => {
  const { signals } = this.state
  const config = {
    iceServers: [
      {
        urls: 'stun:stun1.l.google.com:19302',
      },
    ],
  }

  let peer = new RTCPeerConnection(config)

  peer.addEventListener('icecandidate', ({ candidate }) => {
    if (candidate) {
      signals.sendICECandidate(peerId, candidate)
    }
  })

  peer.addEventListener('track', (track) => {
    peer.srcObject = track.streams[0]
  })

  return peer
})

StreamerCardList.prototype.getPeer((peerId) => {
  const { peers } = this.state
  if (peers[peerId] === undefined) {
    peers[peerId] = this.newPeer(peerId)
    this.setState({ peers })
  }

  return peers[peerId]
})

Object.setPrototypeOf(StreamerCard.prototype, React.Component.prototype);
