import { useEffect, useState } from 'react';

export default function usePeers(local, signals) {
  const [peers, setPeers] = useState(null);

  async function fetchData() {
    const newPeer = (peerId) => {
      const config = {
        iceServers: [
          {
            urls: 'stun:stun1.l.google.com:19302',
          },
        ],
      };

      const peer = new RTCPeerConnection(config);

      peer.addEventListener('icecandidate', ({ candidate }) => {
        if (candidate) {
          signals.sendICECandidate(peerId, candidate);
        }
      });

      return peer;
    };

    const getPeer = (peerId) => {
      if (peers[peerId] === undefined) {
        peers[peerId] = newPeer(peerId);
      }

      return peers[peerId];
    };

    const removePeer = (peerId) => {
      if (peers[peerId] === undefined) {
        return null;
      }

      const peer = peers[peerId];
      delete peers[peerId];

      return peer;
    };

    const onJoin = async (join) => {
      const peer = getPeer(join.peerId);

      local.stream.getTracks().forEach((track) => {
        peer.addTrack(track, local.stream);
      });

      const offer = await peer.createOffer({
        offerToRecieveVideo: 1,
        offerToRecieveAudo: 1,
      });

      await peer.setLocalDescription(offer);
    };

    const onLeave = async (leave) => {
      const peer = removePeer(leave.peerId);
      if (!peer) {
        return;
      }
      peer.close();
      // remove peer from the list
      // videoElem.removeChild(videoElem) in example, we just need to remove it from the peers list
      //
    };

    const onOffer = async (offer) => {
      const peer = getPeer(offer.peerId);
      peer.setRemoteDescription(offer.offer);

      local.stream.getTracks().forEach((track) => {
        peer.addTrack(track, local.stream);
      });

      const answer = await peer.createAnswer();
      peer.setLocalDescription(answer);

      return answer;
    };

    const onAnswer = async (answer) => {
      const peer = getPeer(answer.peerId);
      await peer.setRemoteDescription(answer.answer);
    };

    const onICECandidate = async (candidate) => {
      const peer = getPeer(candidate.peerId);
      peer.addIceCandidate(candidate.candidate);
    };

    const onConnected = () => {};
    const onDisconnected = () => {};

    setPeers({
      peers,
      onConnected,
      onDisconnected,
      onJoin,
      onLeave,
      onOffer,
      onAnswer,
      onICECandidate,
    });
  }

  useEffect(() => {
    fetchData();
  }, [fetchData]);

  return peers;
}
