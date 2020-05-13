

class SignallingServer extends EventTarget {
  constructor(mySecret) {
    super()
    this.websocket = null
  }

  connect(url) {
    let websocket = new WebSocket(url)
    this.websocket = websocket

    websocket.addEventListener("open", (event) => {
      console.log(event)
      this.dispatchEvent(new Event("connected"))
    })
  
    websocket.addEventListener("close", (event) => {
      console.log(event)
      clearInterval(ping)
      this.dispatchEvent(new Event("disconnected"))
    })
  
    websocket.addEventListener("message", (message) => {
      console.log(message)
  
      // TODO consume joins/leaves
      //this.dispatchEvent("peer-join")
      //this.dispatchEvent("peer-leave")
    })
  
    let ping = setInterval(() => {
      websocket.send("heartbeat")
    }, 30000)
  }
}
