package configuration

import "github.com/pion/webrtc/v4"

var (
	// https://w3c.github.io/webrtc-pc/#rtcconfiguration-dictionary
	PeerConnectionConfiguration = webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}
)
