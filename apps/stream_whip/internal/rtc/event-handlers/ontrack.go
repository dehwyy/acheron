package eventhandlers

import (
	"github.com/dehwyy/mugen/apps/stream_whip/internal/rtc/mediastream"
	"github.com/pion/webrtc/v4"
)

func NewOnTrackHandler(mediaStream *mediastream.MediaStream) func(*webrtc.TrackRemote, *webrtc.RTPReceiver) {
	return func(track *webrtc.TrackRemote, recv *webrtc.RTPReceiver) {
		for {
			pkt, _, err := track.ReadRTP()
			if err != nil {
				panic(err)
			}

			switch pkt.PayloadType {
			// H264
			case 96:
				if err = mediaStream.Video.WriteRTP(pkt); err != nil {
					panic(err)
				}
			// Opus
			case 111:

				if err = mediaStream.Audio.WriteRTP(pkt); err != nil {
					panic(err)
				}
			}
		}
	}
}
