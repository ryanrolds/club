export default class SignalingServer extends EventTarget {
  constructor() {
    super()
    this.websocket = null
  }

  connect(url) {
    const websocket = new WebSocket(url)
    this.websocket = websocket

    const ping = setInterval(() => {
      if (websocket.readyState < 1) {
        return
      }

      websocket.send(
        JSON.stringify({ type: 'heartbeat', destId: 'server', payload: {} })
      )
    }, 30000)

    websocket.addEventListener('open', (event) => {
      this.dispatchEvent(new Event('connected', event))
    })

    websocket.addEventListener('close', (event) => {
      clearInterval(ping)
      this.dispatchEvent(new Event('disconnected', event))
    })

    websocket.addEventListener('message', (message) => {
      if (!message.data) {
        console.log('message missing data', message)
        return
      }

      let parsed = null
      try {
        parsed = JSON.parse(message.data)
      } catch (err) {
        console.log('problem parsing message payload', err, message)
        return
      }

      switch (parsed.type) {
        case 'join':
          this.dispatchEvent(
            new CustomEvent('join', {
              detail: { peerId: parsed.peerId, join: parsed.payload },
            })
          )
          break
        case 'leave':
          this.dispatchEvent(
            new CustomEvent('leave', {
              detail: { peerId: parsed.peerId, leave: parsed.payload },
            })
          )
          break
        case 'offer':
          this.dispatchEvent(
            new CustomEvent('offer', {
              detail: { peerId: parsed.peerId, offer: parsed.payload },
            })
          )
          break
        case 'answer':
          this.dispatchEvent(
            new CustomEvent('answer', {
              detail: { peerId: parsed.peerId, answer: parsed.payload },
            })
          )
          break
        case 'icecandidate':
          this.dispatchEvent(
            new CustomEvent('icecandidate', {
              detail: { peerId: parsed.peerId, candidate: parsed.payload },
            })
          )
          break
        default:
          console.log('unknown message type', parsed)
      }
    })
  }

  sendJoin() {
    this.websocket.send(JSON.stringify({ type: 'join', destId: '', payload: {} }))
  }

  sendLeave() {
    this.websocket.send(
      JSON.stringify({ type: 'leave', destId: '', payload: { reason: 'exit' } })
    )
  }

  sendOffer(peerId, offer) {
    this.websocket.send(
      JSON.stringify({ type: 'offer', destId: peerId, payload: offer })
    )
  }

  sendAnswer(peerId, answer) {
    this.websocket.send(
      JSON.stringify({ type: 'answer', destId: peerId, payload: answer })
    )
  }

  sendICECandidate(peerId, candidate) {
    this.websocket.send(
      JSON.stringify({ type: 'icecandidate', destId: peerId, payload: candidate })
    )
  }
}
