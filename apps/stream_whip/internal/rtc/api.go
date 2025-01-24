package rtc

import (
	"errors"

	"github.com/dehwyy/mugen/apps/stream_whip/internal/rtc/configuration"
	"github.com/dehwyy/mugen/apps/stream_whip/internal/rtc/interceptors"
	"github.com/pion/webrtc/v4"
)

type API struct {
	webrtcAPI         *webrtc.API
	peerConfiguration *webrtc.Configuration
}

func NewAPI() (*API, error) {
	engine, mediaEngineErr := newMediaEngine()
	settingEngine, settingEngineErr := newSettingEngine()
	if err := errors.Join(mediaEngineErr, settingEngineErr); err != nil {
		return nil, err
	}

	interceptorRegistry, err := interceptors.New()
	if err != nil {
		return nil, err
	}

	api := webrtc.NewAPI(
		webrtc.WithMediaEngine(engine),
		webrtc.WithSettingEngine(*settingEngine),
		webrtc.WithInterceptorRegistry(interceptorRegistry),
	)

	return &API{
		webrtcAPI:         api,
		peerConfiguration: &configuration.PeerConnectionConfiguration,
	}, nil
}

func (api *API) NewPeerConnection() (*webrtc.PeerConnection, error) {
	return api.webrtcAPI.NewPeerConnection(*api.peerConfiguration)
}
