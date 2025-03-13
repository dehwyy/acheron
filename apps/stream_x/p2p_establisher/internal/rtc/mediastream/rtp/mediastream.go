package rtp

type MediaStreamRTPSender struct {
	AudioRtpSender *RTPSender
	VideoRtpSender *RTPSender
}

func (sender *MediaStreamRTPSender) IsEmpty() bool {
	return sender.AudioRtpSender == nil && sender.VideoRtpSender == nil
}

func (sender *MediaStreamRTPSender) SpawnAckIncomingRTCP() {
	if s := sender.AudioRtpSender; s != nil {
		go s.AckIncomingRTCP()
	}

	if s := sender.VideoRtpSender; s != nil {
		go s.AckIncomingRTCP()
	}
}
