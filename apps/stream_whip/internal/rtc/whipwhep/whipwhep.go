package whipwhep

import (
	"errors"
	"fmt"

	eventhandlers "github.com/dehwyy/acheron/apps/stream_whip/internal/rtc/event-handlers"
	"github.com/dehwyy/acheron/apps/stream_whip/internal/rtc/mediastream"
	"github.com/pion/webrtc/v4"
)

// @Returns:
//   - LocalSDPOffer: string
func exchangeSDPOffers(conn *webrtc.PeerConnection, offer string) (string, error) {
	conn.OnICEConnectionStateChange(eventhandlers.NewOnICEConnectionStateChangeHandler(conn))
	conn.OnICECandidate(eventhandlers.NewOnICECandidateHandler())

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
func HandleWhipConn(conn *webrtc.PeerConnection, streamName, offer string) (string, error) {
	var err error

	if _, err = conn.AddTransceiverFromKind(webrtc.RTPCodecTypeVideo); err != nil {
		return "", err
	}

	if _, err = conn.AddTransceiverFromKind(webrtc.RTPCodecTypeAudio); err != nil {
		return "", err
	}

	mediaStream, err := mediastream.New(streamName, mediastream.AudioVideo)
	if err != nil {
		if errors.Is(err, mediastream.ErrAlreadyExists) {
			return "", fmt.Errorf("%w: token = %s", err, streamName)
		}
		return "", fmt.Errorf("failed to create media stream: %w", err)
	}

	conn.OnTrack(eventhandlers.NewOnTrackHandler(mediaStream))

	return exchangeSDPOffers(conn, offer)
}

// @Returns:
//   - LocalSDPOffer: string
func HandleWhepConn(conn *webrtc.PeerConnection, streamName, offer string) (string, error) {
	mediaStream, err := mediastream.Get(streamName)
	if err != nil {
		if errors.Is(err, mediastream.ErrNotExists) {
			return "", fmt.Errorf("%w: stream = %s", err, streamName)
		}
		return "", fmt.Errorf("failed to get media stream: %w", err)
	}

	rtpSender, err := mediaStream.AddToPeerConnection(conn)
	if err != nil {
		return "", err
	}

	rtpSender.SpawnAckIncomingRTCP()

	return exchangeSDPOffers(conn, offer)
}
