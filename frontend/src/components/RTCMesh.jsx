import React, { Component } from 'react';
import RTCVideo from './RTCVideo.jsx';
import Websocket from './Websocket.jsx';
import PeerConnection from './PeerConnection.jsx';
import { DEFAULT_CONSTRAINTS, DEFAULT_ICE_SERVERS, TYPE_ROOM, TYPE_ANSWER } from './functions/constants';
import { generateRoomKey, createMessage, createPayload } from './functions/utils';

class RTCMesh extends Component {
  constructor(props) {
    super(props);
    const {mediaConstraints, URL } = props;
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

  openCamera = async (fromHandleOffer) => {
    const { mediaConstraints, localMediaStream } = this.state;
    try {
      if (!localMediaStream) {
        let mediaStream;
        if(this.wantCamera) mediaStream = await navigator.mediaDevices.getUserMedia(mediaConstraints);
        else mediaStream = await navigator.mediaDevices.getDisplayMedia(mediaConstraints);

        return fromHandleOffer === true ? mediaStream : this.setState({ localMediaStream: mediaStream });
      }
    } catch(error) {
      console.error('getUserMedia Error: ', error)
    }
  }

  handleOffer = async (data) => {
    const { localMediaStream, roomKey, socketID } = this.state;
    const { payload } = data;
    await this.rtcPeerConnection.setRemoteDescription(payload.message);
    let mediaStream = localMediaStream
    if (!mediaStream) mediaStream = await this.openCamera(true);
    this.setState({ connectionStarted: true, localMediaStream: mediaStream }, async function() {
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

  handleShareDisplay = async() => {
    this.wantCamera = !this.wantCamera
    if(this.state.connectionStarted){
      const { mediaConstraints, localMediaStream } = this.state;
      let mediaStream;
      if(this.wantCamera) mediaStream = await navigator.mediaDevices.getUserMedia(mediaConstraints)
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
      const roomData = createMessage(TYPE_ROOM, createPayload(key, socketID));
      this.setState({ roomKey: key })
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
    const sendMessage = this.socket.send.bind(this.socket);

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
        <PeerConnection
          rtcPeerConnection={this.rtcPeerConnection}
          iceServers={iceServers}
          localMediaStream={localMediaStream}
          addRemoteStream={this.addRemoteStream}
          startConnection={connectionStarted}
          sendMessage={sendMessage}
          roomInfo={{ socketID, roomKey }}
        />
        <RTCVideo mediaStream={localMediaStream} />
        <RTCVideo mediaStream={remoteMediaStream} />

        <section className='button-container'>
          <div className='button button--start-color' onClick={this.openCamera}>
            <button onClick={this.openCamera}>Start Video</button>
          </div>
          <button onClick={this.handleShareDisplay}>Share Screen</button>
          <div className='button button--stop-color' onClick={null}>
          </div>
        </section>
      </>
    );
  }
}
export default RTCMesh;
