package rtc

import (
	"github.com/dehwyy/acheron/apps/stream_x/p2p_establisher/internal/rtc/codecs"
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

func newSettingEngine() (*webrtc.SettingEngine, error) {
	settingEngine := &webrtc.SettingEngine{}

	// mux, err := ice.NewMultiUDPMuxFromPort(8998)
	// if err != nil {
	// 	return nil, err
	// }
	// settingEngine.SetICEUDPMux(mux)

	return settingEngine, nil
}
