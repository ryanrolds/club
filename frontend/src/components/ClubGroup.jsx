import React, { Component } from 'react';
import ClubVideo from './ClubMember.jsx';
import Websocket from './Websocket.jsx';
import PeerConnection from './PeerConnection.jsx';
import { DEFAULT_CONSTRAINTS, DEFAULT_ICE_SERVERS, TYPE_ROOM, TYPE_ANSWER } from './functions/constants';
import { generateRoomKey, createMessage, createPayload } from './functions/utils';

class ClubGroup extends Component {
  constructor(props) {
    super(props);
    const { mediaConstraints, URL } = props;
    // build iceServers config for RTCPeerConnection
    // const iceServerURLs = buildServers(iceServers);
    this.state = {
      iceServers: DEFAULT_ICE_SERVERS,
      mediaConstraints: mediaConstraints || DEFAULT_CONSTRAINTS,
      localMediaStream: null,
      remoteMediaStream: null,
      roomKey: null,
      socketID: null,
      connectionStarted: false,
    };
    this.wantCamera = true;
    this.socket = new WebSocket(URL);
    this.rtcPeerConnection = new RTCPeerConnection({ iceServers: this.state.iceServers });
  }

  handleLocalMedia = async (fromHandleOffer) => {
    const { mediaConstraints, localMediaStream } = this.state;
    try {
      if (!localMediaStream) {
        let mediaStream;
        if (this.wantCamera) mediaStream = await navigator.mediaDevices.getUserMedia(mediaConstraints);
        else mediaStream = await navigator.mediaDevices.getDisplayMedia(mediaConstraints);

        return fromHandleOffer === true ? mediaStream : this.setState({ localMediaStream: mediaStream });
      }
    } catch (error) {
      console.error('getUserMedia Error: ', error)
    }
  }

  handleOffer = async (data) => {
    // Acquire roomkey from req.params.path? lastIndexOf / host?
    // Rename->roomKey->group?
    const { localMediaStream, roomKey, socketID } = this.state;
    debugger
    const { payload } = data;
    // What format is data in right here?
    // Transform data into {  detail: { peerId: parsed.peerId, answer: parsed.payload } }))

    await this.rtcPeerConnection.setRemoteDescription(payload.message);
    let mediaStream = localMediaStream
    if (!mediaStream) mediaStream = await this.handleLocalMedia(true);
    this.setState({ connectionStarted: true, localMediaStream: mediaStream }, async function () {
      const answer = await this.rtcPeerConnection.createAnswer();
      await this.rtcPeerConnection.setLocalDescription(answer);
      const payload = createPayload(roomKey, socketID, answer);
      const answerMessage = createMessage(TYPE_ANSWER, payload);
      this.socket.send(JSON.stringify(answerMessage));
    });
  }

  handleAnswer = async (data) => {
    const { payload } = data;
    await this.rtcPeerConnection.setRemoteDescription(payload.message);
  }

  handleIceCandidate = async (data) => {
    const { message } = data.payload;
    const candidate = JSON.parse(message);
    await this.rtcPeerConnection.addIceCandidate(candidate);
  }

  handleShareDisplay = async () => {
    this.wantCamera = !this.wantCamera
    if (this.state.connectionStarted) {
      const { mediaConstraints, localMediaStream } = this.state;
      let mediaStream;
      if (this.wantCamera) mediaStream = await navigator.mediaDevices.getUserMedia(mediaConstraints)
      else mediaStream = await navigator.mediaDevices.getDisplayMedia(mediaConstraints)

      let screenStream = mediaStream.getVideoTracks()[0]
      const transceiver = this.rtcPeerConnection.getTransceivers()[0]
      localMediaStream.removeTrack(localMediaStream.getTracks()[0])
      localMediaStream.addTrack(screenStream)
      transceiver['sender'].replaceTrack(screenStream)
    }
  }

  sendRoomKey = () => {
    const { roomKey, socketID } = this.state;
    if (!roomKey) {
      const key = generateRoomKey();
      // /room
      const roomData = createMessage(TYPE_ROOM, createPayload(key, socketID));
      this.setState({ roomKey: key, socketID: 1 })
      this.socket.send(JSON.stringify(roomData));
      alert(key);
    }
  }

  handleSocketConnection = (socketID) => {
    this.setState({ socketID });
  }

  handleConnectionReady = (message) => {
    console.log('Inside handleConnectionReady: ', message);
    if (message.startConnection) {
      this.setState({ connectionStarted: message.startConnection });
    }
  }

  addRemoteStream = (remoteMediaStream) => {
    this.setState({ remoteMediaStream });
  }


  render() {
    const {
      localMediaStream,
      remoteMediaStream,
      roomKey,
      socketID,
      iceServers,
      connectionStarted,
    } = this.state;
    this.handleLocalMedia()
    console.log('Ready to send video')
    const sendMessage = this.socket.send.bind(this.socket);
    console.log('Socket ReadyState: ', this.socket.readyState)
    if (this.socket && this.socket.readyState >= 1) this.sendRoomKey()
    console.log(this.socket)
    return (
      <>
        <Websocket
          socket={this.socket}
          setSendMethod={this.setSendMethod}
          handleSocketConnection={this.handleSocketConnection}
          handleConnectionReady={this.handleConnectionReady}
          handleOffer={this.handleOffer}
          handleAnswer={this.handleAnswer}
          handleIceCandidate={this.handleIceCandidate}
        />
        {/* TODO
        * Send Call Offer to Browser 2 (Peer)
        * Verify remoteMediaStream exists and is applied to ClubVideo id=remote
        *
        *
        *
        *
        */}
        {
          localMediaStream ?
            (<>
              <PeerConnection
                rtcPeerConnection={this.rtcPeerConnection}
                iceServers={iceServers}
                localMediaStream={localMediaStream}
                addRemoteStream={this.addRemoteStream}
                startConnection={connectionStarted}
                sendMessage={sendMessage}
                roomInfo={{ socketID, roomKey }}
              />
              <ClubVideo id="local" mediaStream={localMediaStream} muted={true} />
            </>
            ) : null
        }
        {
          remoteMediaStream ?
            (<>
              <ClubVideo id="remote" mediaStream={remoteMediaStream} muted={false} />
            </>) : null
        }
      </>
    );
  }
}
export default ClubGroup;
