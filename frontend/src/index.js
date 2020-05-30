import React from 'react';
import ReactDOM from 'react-dom';
import RTCMesh from './components/RTCMesh.jsx';

let isHTTPS = window.location.protocol !== 'https:'
const wsServer = (isHTTPS ? "ws" : "wss") + "://localhost:3001/room"
ReactDOM.render(<RTCMesh URL={wsServer} />, document.getElementById('root'));
