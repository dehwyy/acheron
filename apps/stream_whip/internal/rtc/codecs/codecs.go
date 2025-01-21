package codecs

import "github.com/pion/webrtc/v4"

const (
	AudioOpusID = 111
	VideoH264ID = 96
)

var (
	PresetAudioOpus = webrtc.RTPCodecParameters{
		RTPCodecCapability: webrtc.RTPCodecCapability{
			MimeType: webrtc.MimeTypeOpus, ClockRate: 48000, Channels: 2, SDPFmtpLine: "", RTCPFeedback: nil,
		},
		PayloadType: AudioOpusID,
	}

	PresetVideoH264 = webrtc.RTPCodecParameters{
		RTPCodecCapability: webrtc.RTPCodecCapability{
			MimeType: webrtc.MimeTypeH264, ClockRate: 90000, Channels: 0, SDPFmtpLine: "", RTCPFeedback: nil,
		},
		PayloadType: VideoH264ID,
	}
)
