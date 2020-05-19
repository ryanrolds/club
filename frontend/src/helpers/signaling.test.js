import React from 'react'
import '@testing-library/jest-dom/extend-expect'
import WS from 'jest-websocket-mock'
import SignalingServer from './signaling'

describe('Messages', () => {
  test('test: JOIN', async () => {
    const server = new WS('ws://localhost:3000', { jsonProtocol: true })
    //const client = new SignalingServer('ws://localhost:3000')
    const client = new WebSocket('ws://localhost:3000')
    await server.connected
    server.send({ type: 'GREETING', payload: 'hello' })
    client.send(JSON.stringify({ type: 'heartbeat', destId: 'server', payload: {} }))
    expect(client.readyState)
  })
})
