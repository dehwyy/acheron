package tracks

import (
	"errors"

	"github.com/pion/webrtc/v4"
)

const (
	rtcpBufSize = 1500
)

type RTPSender struct {
	innerSender *webrtc.RTPSender
}

func (sender *RTPSender) ReadIncoming() {
	rtcpBuf := make([]byte, rtcpBufSize)
	for {
		if _, _, rtcpErr := sender.innerSender.Read(rtcpBuf); rtcpErr != nil {
			return
		}
	}
}

func AddAudioVideoTracks(conn *webrtc.PeerConnection, streamToken string) (*RTPSender, *RTPSender, error) {
	audioTrack, videoTrack, err := getTracks(streamToken)
	if err != nil {
		return nil, nil, err
	}

	audioSender, audioErr := conn.AddTrack(audioTrack)
	videoSender, videoErr := conn.AddTrack(videoTrack)
	if err := errors.Join(audioErr, videoErr); err != nil {
		return nil, nil, nil
	}

	return &RTPSender{audioSender}, &RTPSender{videoSender}, nil
}

// func AddAudioTrack(conn *webrtc.PeerConnection, streamToken string) (*RTPSender, error) {
// 	return addTrack(conn, streamToken, Audio)
// }

// func AddVideoTrack(conn *webrtc.PeerConnection, streamToken string) (*RTPSender, error) {
// 	return addTrack(conn, streamToken, Video)
// }

// func AddAudioVideoTracks(conn *webrtc.PeerConnection, streamToken string) (*RTPSender, *RTPSender, error) {
// 	audioTrack, audioErr := AddAudioTrack(conn, streamToken)
// 	videoTrack, videoErr := AddVideoTrack(conn, streamToken)
// 	if err := errors.Join(audioErr, videoErr); err != nil {
// 		return nil, nil, nil
// 	}

// 	return audioTrack, videoTrack, nil
// }
