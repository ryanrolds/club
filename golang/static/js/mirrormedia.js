
// eslint-disable-next-line no-unused-vars
class MirrorMedia {
  constructor(video, canvas) {
    this.id = video.id
    this.video = video
    this.canvas = canvas
    this.stream = null
  }

  async setupMedia() {
    const stream = this.canvas.captureStream(25)

    this.stream = stream

    this.video.srcObject = stream
  }

  getStream() {
    return this.stream
  }

  async onConnected() {}
  async onDisconnected() {}
}
