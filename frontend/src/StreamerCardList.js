import React, { useState, useEffect } from 'react'
import { makeStyles, Container, Grid } from '@material-ui/core'
import StreamerCard from './StreamerCard'
import SignalingServer from './helpers/signaling'

export default function StreamerCardList() {
  const [peers, setPeers] = useState({ peers: {} })
  const [stream, setStream] = useState({ stream: null })
  const signals = new SignalingServer()

  const classes = makeStyles((theme) => ({
    cardGrid: {
      paddingTop: theme.spacing(1),
      paddingBottom: theme.spacing(1),
    },
  }))

  function newPeer(peerId) {
    const config = {
      iceServers: [
        {
          urls: 'stun:stun1.l.google.com:19302',
        },
      ],
    }

    const peer = new RTCPeerConnection(config)

    peer.addEventListener('icecandidate', ({ candidate }) => {
      if (candidate) {
        signals.sendICECandidate(peerId, candidate)
      }
    })

    peer.addEventListener('track', (track) => {
      const srcObject = track.streams[0]
      peer.srcObject = srcObject
    })

    return peer
  }

  function getPeer(peerId) {
    if (peers[peerId] === undefined) {
      peers[peerId] = newPeer(peerId)
      setPeers({ peers })
    }

    return peers[peerId]
  }

  async function onICECandidate(candidate) {
    const peer = getPeer(candidate.peerId)
    peer.addIceCandidate(candidate.candidate)
  }

  async function onConnected() {
    return false
  }

  async function onDisconnected() {
    return false
  }

  async function onJoin(join) {
    const peer = getPeer(join.peerId)

    stream.getTracks().forEach((track) => {
      peer.addTrack(track, stream)
    })

    const newOffer = await peer.createOffer({
      offerToReceiveVideo: 1,
      offerToReceiveAudio: 1,
    })

    await peer.setLocalDescription(newOffer)

    return newOffer
  }

  async function onOffer(offer) {
    const peer = getPeer(offer.peerId)
    peer.setRemoteDescription(offer.offer)

    stream.getTracks().forEach((track) => {
      peer.addTrack(track, stream)
    })

    const answer = await peer.createAnswer()
    peer.setLocalDescription(answer)

    return answer
  }

  async function onAnswer(answer) {
    const peer = getPeer(answer.peerId)
    await peer.setRemoteDescription(answer.answer)
  }

  async function onLeave() {
    return false
  }

  async function setupSignalEventHandlers() {
    signals.addEventListener('connected', async (event) => {
      console.log('connected to signalling server', event)
      await onConnected()
      // await peers.onConnected()

      signals.sendJoin()
    })

    signals.addEventListener('join', async (event) => {
      console.log('got join', event.join)
      const newOffer = await onJoin(event.detail)
      signals.sendOffer(event.detail.peerId, newOffer)
    })

    signals.addEventListener('offer', async (event) => {
      console.log('got offer', event)
      const answer = await onOffer(event.detail)
      signals.sendAnswer(event.detail.peerId, answer)
    })

    signals.addEventListener('icecandidate', async (event) => {
      console.log('got ice candidate', event)
      onICECandidate(event.detail)
    })

    signals.addEventListener('answer', async (event) => {
      console.log('got answer', event)
      await onAnswer(event.detail)
    })

    signals.addEventListener('leave', async (event) => {
      console.log('got leave', event)
      await onLeave(event.detail)
    })

    signals.addEventListener('disconnected', async (event) => {
      console.log('disconnected from signalling server', event)

      await onDisconnected()
      // await onDisconnected()
    })

    const isHTTPS = window.location.protocol !== 'https:'
    if (!isHTTPS) signals.connect(`ws://${window.location.host}/room`)
    signals.connect(`wss://${window.location.host}/room`)
  }

  useEffect(async () => {
    const opts = { audio: true, video: true }
    const newStream = await navigator.mediaDevices.getUserMedia(opts)
    setStream({ newStream })
    setupSignalEventHandlers()
  })

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
}
