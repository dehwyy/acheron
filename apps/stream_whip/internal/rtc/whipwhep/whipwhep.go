package whipwhep

import (
	eventhandlers "github.com/dehwyy/mugen/apps/stream_whip/internal/rtc/event-handlers"
	"github.com/dehwyy/mugen/apps/stream_whip/internal/rtc/tracks"
	"github.com/pion/webrtc/v4"
)

// @Returns:
//   - LocalSDPOffer: string
func exchangeSDPOffers(conn *webrtc.PeerConnection, offer string) (string, error) {
	conn.OnICEConnectionStateChange(eventhandlers.NewOnICEConnectionStateChangeHandler(conn))

	if err := conn.SetRemoteDescription(webrtc.SessionDescription{
		Type: webrtc.SDPTypeOffer,
		SDP:  offer,
	}); err != nil {
		return "", err
	}

	gatherComplete := webrtc.GatheringCompletePromise(conn)

	answer, err := conn.CreateAnswer(&webrtc.AnswerOptions{})
	if err != nil {
		return "", err
	}

	if err = conn.SetLocalDescription(answer); err != nil {
		return "", err
	}

	<-gatherComplete

	return conn.LocalDescription().SDP, nil
}

// @Returns:
//   - LocalSDPOffer: string
func HandleWhipConn(conn *webrtc.PeerConnection, streamToken, offer string) (string, error) {
	var err error

	if _, err = conn.AddTransceiverFromKind(webrtc.RTPCodecTypeVideo); err != nil {
		return "", err
	}

	if _, err = conn.AddTransceiverFromKind(webrtc.RTPCodecTypeAudio); err != nil {
		return "", err
	}

	audioTrack, videoTrack, err := tracks.NewDefaultAudioVideoTracks(streamToken)
	if err != nil {
		return "", err
	}

	conn.OnTrack(eventhandlers.NewOnTrackHandler(audioTrack, videoTrack))

	return exchangeSDPOffers(conn, offer)
}

// @Returns:
//   - LocalSDPOffer: string
func HandleWhepConn(conn *webrtc.PeerConnection, streamToken, offer string) (string, error) {
	audioRtpSender, videoRtpSender, err := tracks.AddAudioVideoTracks(conn, streamToken)
	if err != nil {
		return "", nil
	}

	go audioRtpSender.ReadIncoming()
	go videoRtpSender.ReadIncoming()

	return exchangeSDPOffers(conn, offer)
}
