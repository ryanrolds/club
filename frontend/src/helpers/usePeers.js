import { useEffect, useState } from 'react'

export default function usePeers() {
  const [peers, setPeers] = useState(null)

  async function fetchData() {
    return setPeers([{ id: '1', stream: '' }, { id: '2', stream: '' }])
  }

  useEffect(() => { fetchData() }, [])

  return peers
}
