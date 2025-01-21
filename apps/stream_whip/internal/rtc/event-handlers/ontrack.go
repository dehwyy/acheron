package eventhandlers

import (
	"github.com/pion/webrtc/v4"
)

func NewOnTrackHandler(audioTrack, videoTrack *webrtc.TrackLocalStaticRTP) func(*webrtc.TrackRemote, *webrtc.RTPReceiver) {
	return func(track *webrtc.TrackRemote, recv *webrtc.RTPReceiver) {
		for {
			pkt, _, err := track.ReadRTP()
			if err != nil {
				panic(err)
			}

			switch pkt.PayloadType {
			// H264
			case 96:

				if err = videoTrack.WriteRTP(pkt); err != nil {
					panic(err)
				}
			// Opus
			case 111:

				if err = audioTrack.WriteRTP(pkt); err != nil {
					panic(err)
				}
			}
		}
	}
}
