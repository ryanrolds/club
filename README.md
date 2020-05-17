# Club

![Travis Build Status](https://travis-ci.org/ryanrolds/club.svg?branch=master)

WebRTC video chat application written in JS and Go.

## Todos

* ~~~Propagate "leaves" and update client to remove peers/videos that left~~~
* See if there are audio options that can be implemented
* Create grid UI using Material UI https://github.com/ryanrolds/club/issues/1
* Add mute buttons for self and other users
* Add video off button for self
* Add leave button and join button
* Implement multiple rooms
* Implement join password for rooms
* Create UI for providing room ID and password (if passworded room)
* Get ICE Server(s) from env var
* Get ICE Servers (STUN/TURN) from successful join response (don't store in client)

## Setup

Requires Go 1.14+ and Node 12+.

```
make install
```

## Running

Each service can be started seperately with `make run` in the `frontend` and `golang` directories. Otherwise, from the root both services can be started using Docker Compose:

```
make run
```

For extra debugging information use `make run-debug` in the service's directory.

Open http://localhost:3000 for React Client and http://localhost:3001 for the signaling server and examples.
