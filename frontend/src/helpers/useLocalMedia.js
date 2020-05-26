import { useEffect, useState } from "react"

export default function useLocalMedia(){
  const [localMedia, setLocalMedia] = useState(null)

  async function fetchData(){
    const stream = await navigator.mediaDevices.getUserMedia({ audio: true, video: true })
    setLocalMedia({id: stream.id, stream})
  }

  useEffect(()=> {fetchData()},[])

  return localMedia
}
