package rtc

import (
	"github.com/dehwyy/mugen/apps/stream_whip/internal/rtc/codecs"
	"github.com/pion/webrtc/v4"
)

func newMediaEngine() (*webrtc.MediaEngine, error) {
	mediaEngine := &webrtc.MediaEngine{}

	if err := mediaEngine.RegisterCodec(codecs.PresetAudioOpus, webrtc.RTPCodecTypeAudio); err != nil {
		return nil, err
	}

	if err := mediaEngine.RegisterCodec(codecs.PresetVideoH264, webrtc.RTPCodecTypeVideo); err != nil {
		return nil, err
	}

	return mediaEngine, nil
}
