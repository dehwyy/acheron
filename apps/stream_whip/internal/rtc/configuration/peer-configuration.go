package configuration

import "github.com/pion/webrtc/v4"

var (
	PeerConnectionConfiguration = webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}
)
