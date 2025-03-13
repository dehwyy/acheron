package rtp

import "github.com/pion/webrtc/v4"

const (
	rtcpBufSize = 1500
)

type RTPSender struct {
	inner *webrtc.RTPSender
}

func NewRTPSender(rtpSender *webrtc.RTPSender) *RTPSender {
	return &RTPSender{rtpSender}
}

func (sender *RTPSender) AckIncomingRTCP() {
	rtcpBuf := make([]byte, rtcpBufSize)
	for {
		if _, _, rtcpErr := sender.inner.Read(rtcpBuf); rtcpErr != nil {
			return
		}
	}
}
