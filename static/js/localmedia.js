
class LocalMedia {
  constructor(id) {
    this.id = id
    this.stream = null
  }

  async setupMedia() {
    const opts = {audio: true, video: true}
    const stream = await navigator.mediaDevices.getUserMedia(opts)
    this.stream = stream

    const localMedia = document.querySelector(this.id)
    localMedia.srcObject = stream
  }
}
