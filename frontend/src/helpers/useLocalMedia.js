import { useEffect, useState } from 'react'

export default function useLocalMedia() {
  const [localMedia, setLocalMedia] = useState(null)

  async function fetchData() {
    const stream = await navigator.mediaDevices.getUserMedia({
      audio: true,
      video: true,
    })
    const onConnected = () => {}
    const onDisconnected = () => {}
    setLocalMedia({
      id: stream.id,
      stream,
      muted: true,
      onConnected,
      onDisconnected,
    })
  }

  useEffect(() => {
    fetchData()
  }, [])

  return localMedia
}
