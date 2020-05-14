
class LocalMedia {
  constructor(id) {
    this.id = id
    this.stream = null
  }

  async setupMedia() {
    const opts = {audio: true, video: true}
    const stream = await navigator.mediaDevices.getUserMedia(opts)
    this.stream = stream

    const videoElm = document.querySelector(`${this.id} video`)
    videoElm.srcObject = stream
  }

  getStream() {
    return this.stream
  }

  async onConnected() {}
  async onDisconnected() {}
}
