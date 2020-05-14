# Club

WebRTC video chat application written in JS and Go.

## Todos

* Propagate "leaves" and update client to remove peers/videos that left
* See if there are audio options that can be implemented
* Create grid UI using Material UI
* Add mute buttons for self and other users
* Add video off button for self
* Add leave button and join button
* Implement join password
* Get ICE Server(s) from env var
* Get ICE Servers (STUN/TURN) from successful join response (don't store in client)

## Setup

Requires Go 1.14+.

## Running

```
go build
./club
```

Open http://localhost:3000 in your browser.
