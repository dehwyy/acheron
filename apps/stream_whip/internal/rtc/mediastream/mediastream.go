package mediastream

import (
	"errors"

	"github.com/dehwyy/mugen/apps/stream_whip/internal/rtc/mediastream/rtp"
	"github.com/dehwyy/mugen/apps/stream_whip/internal/rtc/mediastream/track"
	"github.com/pion/webrtc/v4"
)

type MediaStreamKind = int

const (
	Audio      MediaStreamKind = 0b01
	Video      MediaStreamKind = 0b10
	AudioVideo MediaStreamKind = 0b11
)

type MediaStream struct {
	Audio *webrtc.TrackLocalStaticRTP
	Video *webrtc.TrackLocalStaticRTP
}

var (
	CreatedStreams = map[string]*MediaStream{}

	ErrAlreadyExists = errors.New("media stream already exists")
	ErrNotExists     = errors.New("media stream not exists yet")
	ErrEmpty         = errors.New("media stream is empty")
)

func New(token string, kind MediaStreamKind) (*MediaStream, error) {
	if _, ok := CreatedStreams[token]; ok {
		return nil, ErrAlreadyExists
	}

	stream := &MediaStream{}

	if (kind & Audio) != 0 {
		audioTrack, err := track.NewAudioOpusTrack(token)
		if err != nil {
			return nil, err
		}

		stream.Audio = audioTrack
	}

	if (kind & Video) != 0 {
		videoTrack, err := track.NewVideoH264Track(token)
		if err != nil {
			return nil, err
		}

		stream.Video = videoTrack
	}

	CreatedStreams[token] = stream

	return stream, nil
}

func Get(token string) (*MediaStream, error) {
	if stream, ok := CreatedStreams[token]; ok {
		return stream, nil
	}

	return nil, ErrNotExists
}

func (s *MediaStream) AddToPeerConnection(conn *webrtc.PeerConnection) (*rtp.MediaStreamRTPSender, error) {
	rtpSender := &rtp.MediaStreamRTPSender{}

	if s.Audio != nil {
		audioRtpSender, err := conn.AddTrack(s.Audio)
		if err != nil {
			return nil, err
		}

		rtpSender.AudioRtpSender = rtp.NewRTPSender(audioRtpSender)
	}

	if s.Video != nil {
		videoRtpSender, err := conn.AddTrack(s.Video)
		if err != nil {
			return nil, err
		}

		rtpSender.VideoRtpSender = rtp.NewRTPSender(videoRtpSender)
	}

	if rtpSender.IsEmpty() {
		return nil, ErrEmpty
	}

	return rtpSender, nil
}
