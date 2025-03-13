package track

import (
	"github.com/dehwyy/acheron/apps/stream_x/p2p_establisher/internal/rtc/codecs"
	"github.com/pion/webrtc/v4"
)

func NewAudioOpusTrack(streamID string) (*webrtc.TrackLocalStaticRTP, error) {
	return webrtc.NewTrackLocalStaticRTP(codecs.PresetAudioOpus.RTPCodecCapability, "audio", streamID)
}

func NewVideoH264Track(streamID string) (*webrtc.TrackLocalStaticRTP, error) {
	return webrtc.NewTrackLocalStaticRTP(codecs.PresetVideoH264.RTPCodecCapability, "video", streamID)
}
