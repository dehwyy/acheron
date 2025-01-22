package track

import (
	"github.com/dehwyy/mugen/apps/stream_whip/internal/rtc/codecs"
	"github.com/pion/webrtc/v4"
)

func NewAudioOpusTrack(streamID string) (*webrtc.TrackLocalStaticRTP, error) {
	return webrtc.NewTrackLocalStaticRTP(codecs.PresetAudioOpus.RTPCodecCapability, "audio", streamID)
}

func NewVideoH264Track(streamID string) (*webrtc.TrackLocalStaticRTP, error) {
	return webrtc.NewTrackLocalStaticRTP(codecs.PresetVideoH264.RTPCodecCapability, "video", streamID)
}
