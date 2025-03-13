package eventhandlers

import "github.com/pion/webrtc/v4"

func NewOnICEConnectionStateChangeHandler(conn *webrtc.PeerConnection) func(webrtc.ICEConnectionState) {
	return func(state webrtc.ICEConnectionState) {
		if state == webrtc.ICEConnectionStateFailed {
			_ = conn.Close()
		}
	}
}
