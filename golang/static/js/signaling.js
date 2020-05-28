// eslint-disable-next-line no-unused-vars
class SignalingServer extends EventTarget {
  constructor() {
    super()
    this.websocket = null
  }

  connect(url) {
    let websocket = new WebSocket(url)
    this.websocket = websocket

    websocket.addEventListener("open", (event) => {
      console.log(`connected: ${event}`)
      this.dispatchEvent(new Event("connected"))
    })

    websocket.addEventListener("close", (event) => {
      clearInterval(ping)
      console.log(`closed: ${event}`)
      this.dispatchEvent(new Event("disconnected"))
    })

    websocket.addEventListener("message", (message) => {
      if (!message.data) {
        console.log("message missing data", message)
        return
      }

      let parsed = null
      try {
        parsed = JSON.parse(message.data)
      } catch (err) {
        console.log("problem parsing message payload", err, message)
        return
      }

      switch (parsed.type) {
        case "join":
          this.dispatchEvent(new CustomEvent("join", { detail: { peerId: parsed.peerId, join: parsed.payload } }))
          break
        case "leave":
          this.dispatchEvent(new CustomEvent("leave", { detail: { peerId: parsed.peerId, leave: parsed.payload } }))
          break
        case "offer":
          this.dispatchEvent(new CustomEvent("offer", { detail: { peerId: parsed.peerId, offer: parsed.payload } }))
          break
        case "answer":
          this.dispatchEvent(new CustomEvent("answer", { detail: { peerId: parsed.peerId, answer: parsed.payload } }))
          break
        case "icecandidate":
          this.dispatchEvent(new CustomEvent("icecandidate", { detail: { peerId: parsed.peerId, candidate: parsed.payload } }))
          break
        case "error":
          console.log("Error processing message:", parsed.payload.error)
          console.log(parsed)
          break
        default:
          console.log("unknown message type", parsed)
      }
    })

    let ping = setInterval(() => {
      websocket.send(JSON.stringify({ type: "heartbeat", destId: "server", payload: {} }))
    }, 30000)
  }

  sendJoin(group) {
    this.websocket.send(JSON.stringify({ type: "join", destId: "", payload: { group }}))
  }

  sendLeave() {
    this.websocket.send(JSON.stringify({ type: "leave", destId: "", payload: { "reason": "exit" } }))
  }

  sendOffer(peerId, offer) {
    this.websocket.send(JSON.stringify({ type: "offer", destId: peerId, payload: offer }))
  }

  sendAnswer(peerId, answer) {
    this.websocket.send(JSON.stringify({ type: "answer", destId: peerId, payload: answer }))
  }

  sendICECandidate(peerId, candidate) {
    this.websocket.send(JSON.stringify({ type: "icecandidate", destId: peerId, payload: candidate }))
  }
}
