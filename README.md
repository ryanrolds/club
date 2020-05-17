# Club

![Travis Build Status](https://travis-ci.org/ryanrolds/club.svg?branch=master)

WebRTC video chat application written in JS and Go.

## Todos

* [DONE] Propagate "leaves" and update client to remove peers/videos that left
* See if there are audio options that can be implemented - https://github.com/ryanrolds/club/issues/10
* Create grid UI using Material UI - https://github.com/ryanrolds/club/issues/1
  * Add mute buttons for self and other users
  * Add video off button for self
  * Add leave button and join button
* Implement multiple rooms
* Implement join password for rooms
* Decide if we should switch WebSocket and server interactions to channels
* Create UI for providing room ID and password (if passworded room)
* Get ICE Server(s) from env var
* Get ICE Servers (STUN/TURN) from successful join response (don't store in client)

## Setup

Requires Go 1.14+.

```
make install
```

## Running

```
make run
```

For extra debugging information use `make run-debug`

Open http://localhost:3000 in your browser.
