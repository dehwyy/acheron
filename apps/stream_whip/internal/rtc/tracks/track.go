package tracks

import (
	"errors"
	"fmt"

	"github.com/dehwyy/mugen/apps/stream_whip/internal/rtc/codecs"
	"github.com/pion/webrtc/v4"
)

type TrackKind = string

var (
	createdTracks = map[string]struct {
		audio webrtc.TrackLocal
		video webrtc.TrackLocal
	}{}
)

const (
	Audio TrackKind = "audio"
	Video TrackKind = "video"
)

func formatStreamID(streamToken string) string {
	return streamToken
}

func registerNewTrack(streamID string, trackKind TrackKind, track webrtc.TrackLocal) {
	value, exists := createdTracks[streamID]
	if !exists {
		value = struct {
			audio webrtc.TrackLocal
			video webrtc.TrackLocal
		}{}
	}

	switch trackKind {
	case Audio:
		value.audio = track
	case Video:
		value.video = track
	}

	createdTracks[streamID] = value
}

func getTracks(streamToken string) (webrtc.TrackLocal, webrtc.TrackLocal, error) {
	tracks, ok := createdTracks[formatStreamID(streamToken)]
	if !ok {
		return nil, nil, fmt.Errorf("tracks not found for streamToken <%s>", streamToken)
	}
	return tracks.audio, tracks.video, nil
}

func NewAudioOpusTrack(streamToken string) (*webrtc.TrackLocalStaticRTP, error) {
	trackStreamID := formatStreamID(streamToken)

	track, err := webrtc.NewTrackLocalStaticRTP(codecs.PresetAudioOpus.RTPCodecCapability, Audio, trackStreamID)
	if err != nil {
		return nil, err
	}

	registerNewTrack(trackStreamID, Audio, track)

	return track, nil
}

func NewVideoH264Track(streamToken string) (*webrtc.TrackLocalStaticRTP, error) {
	trackStreamID := formatStreamID(streamToken)

	track, err := webrtc.NewTrackLocalStaticRTP(codecs.PresetVideoH264.RTPCodecCapability, Video, trackStreamID)
	if err != nil {
		return nil, err
	}

	registerNewTrack(trackStreamID, Video, track)

	return track, nil
}

func NewDefaultAudioVideoTracks(streamToken string) (*webrtc.TrackLocalStaticRTP, *webrtc.TrackLocalStaticRTP, error) {
	audioTrack, audioErr := NewAudioOpusTrack(streamToken)
	videoTrack, videoErr := NewVideoH264Track(streamToken)

	if err := errors.Join(audioErr, videoErr); err != nil {
		return nil, nil, err
	}

	return audioTrack, videoTrack, nil
}

// func NewTrackByKind(streamToken string, trackKind TrackKind) (*webrtc.TrackLocalStaticRTP, error) {
// 	switch trackKind {
// 	case Audio:
// 		return NewAudioOpusTrack(streamToken)
// 	case Video:
// 		return NewVideoH264Track(streamToken)
// 	}

// 	return nil, nil
// }
