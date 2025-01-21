package rtc

import (
	"github.com/pion/webrtc/v4"
)

type API struct {
	webrtcAPI         *webrtc.API
	peerConfiguration *webrtc.Configuration
}

func NewAPI() (*API, error) {
	engine, err := newMediaEngine()
	if err != nil {
		return nil, err
	}

	api := webrtc.NewAPI(webrtc.WithMediaEngine(engine))

	return &API{
		webrtcAPI:         api,
		peerConfiguration: &peerConnectionConfiguration,
	}, nil
}

func (api *API) NewPeerConnection() (*webrtc.PeerConnection, error) {
	return api.webrtcAPI.NewPeerConnection(*api.peerConfiguration)
}
