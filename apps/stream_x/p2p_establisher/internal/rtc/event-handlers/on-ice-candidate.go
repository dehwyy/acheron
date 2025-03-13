package eventhandlers

import "github.com/pion/webrtc/v4"

func NewOnICECandidateHandler() func(*webrtc.ICECandidate) {
	return func(_ *webrtc.ICECandidate) {
	}
}
