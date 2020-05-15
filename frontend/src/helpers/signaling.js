

export default class SignalingServer extends EventTarget {
  constructor(mySecret) {
    super()
    this.websocket = null
  }

  connect(url) {
    let websocket = new WebSocket(url)
    this.websocket = websocket

    websocket.addEventListener("open", (event) => {
      console.log("got open", event)
      this.dispatchEvent(new Event("connected"))
    })

    websocket.addEventListener("close", (event) => {
      console.log("got close", event)
      clearInterval(ping)
      this.dispatchEvent(new Event("disconnected"))
    })

    websocket.addEventListener("message", (message) => {
      console.log("received message", message)

      if (!message.data) {
        return
      }

      let parsed = null
      try {
        parsed = JSON.parse(message.data)
      } catch(err) {
        console.log("problem parsing message payload", err, message)
        return
      }

      switch(parsed.type) {
        case "join":
          this.dispatchEvent(new CustomEvent("join", {detail: {peerId: parsed.peerId, join: parsed.payload}}))
          break;
        case "offer":
          this.dispatchEvent(new CustomEvent("offer",  {detail: {peerId: parsed.peerId,offer: parsed.payload}}))
          break;
        case "answer":
          this.dispatchEvent(new CustomEvent("answer", {detail: {peerId: parsed.peerId, answer: parsed.payload}}))
          break;
        case "icecandidate":
          this.dispatchEvent(new CustomEvent("icecandidate", {detail: {peerId: parsed.peerId, candidate: parsed.payload}}))
          break;
        default:
          console.log("unknown message type", parsed)
      }
    })

    let ping = setInterval(() => {
      websocket.send(JSON.stringify({type: "heartbeat", destId: "server", payload: {}}))
    }, 30000)
  }

  sendJoin() {
    this.websocket.send(JSON.stringify({type: "join", destId: "", payload: {}}))
  }

  sendOffer(peerId, offer) {
    this.websocket.send(JSON.stringify({type: "offer", destId: peerId, payload: offer}))
  }

  sendAnswer(peerId, answer) {
    this.websocket.send(JSON.stringify({type: "answer", destId: peerId, payload: answer}))
  }

  sendICECandidate(peerId, candidate) {
    this.websocket.send(JSON.stringify({type: "icecandidate",  destId: peerId, payload: candidate}))
  }
}
