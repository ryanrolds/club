import React from 'react'
import { makeStyles, Container, Grid } from '@material-ui/core'
import StreamerCard from './StreamerCard'
import SignalingServer from './helpers/signaling'

export default class StreamerCardList extends React.Component {
  /* ESLint won't support class field declarations until TC39
  * hits stage 4. https://github.com/tc39/proposal-class-fields
  */
  state = {
    peers: {},
    signals: new SignalingServer(),
    stream: null,
    offer: null,
  }

  useStyles() {
    return makeStyles((theme) => ({
      cardGrid: {
        paddingTop: theme.spacing(1),
        paddingBottom: theme.spacing(1),
      },
    }))
  }

  render() {
    const classes = this.useStyles()

    return (
      <Container className={classes.cardGrid}>
        <Grid container spacing={1}>
          <Grid item key={this.state.stream} xs={12} sm={6}>
            <StreamerCard stream={this.state.stream} muted />
          </Grid>
          {this.state.peers.length &&
            this.state.peers.map((peer) => (
              <Grid item key={peer} xs={12} sm={6}>
                <StreamerCard stream={peer.srcObject} />
              </Grid>
            ))}
        </Grid>
      </Container>
    )
  }

  componentDidMount() {
    this.setupMedia()
    this.setupSignalEventHandlers()
  }

  async onConnected() {}

  async onDisconnected() {}

  async onJoin(join) {
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
  }

  async onOffer(offer) {
    const { stream } = this.state
    let peer = this.getPeer(offer.peerId)
    peer.setRemoteDescription(offer.offer)

    stream.getTracks().forEach((track) => {
      peer.addTrack(track, stream)
    })

    const answer = await peer.createAnswer()
    peer.setLocalDescription(answer)

    return answer
  }

  async onAnswer(answer) {
    let peer = this.getPeer(answer.peerId)
    await peer.setRemoteDescription(answer.answer)
  }

  async onICECandidate(candidate) {
    let peer = this.getPeer(candidate.peerId)
    peer.addIceCandidate(candidate.candidate)
  }

  newPeer(peerId) {
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
  }

  getPeer(peerId) {
    const { peers } = this.state
    if (peers[peerId] === undefined) {
      peers[peerId] = this.newPeer(peerId)
      this.setState({ peers })
    }

    return peers[peerId]
  }

  async setupMedia() {
    const opts = { audio: true, video: true }
    const stream = await navigator.mediaDevices.getUserMedia(opts)
    this.setState({ stream })
  }

  async setupSignalEventHandlers() {
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
  }
}
