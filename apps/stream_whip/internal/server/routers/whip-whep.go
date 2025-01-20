package routers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/dehwyy/mugen/libraries/go/logg"
	"github.com/gin-gonic/gin"
	"github.com/pion/interceptor"
	"github.com/pion/interceptor/pkg/intervalpli"
	"github.com/pion/webrtc/v4"
)

const (
	routeHandleWhip = "/whip"
	routeHandleWhep = "/whep"
)

var (
	videoTrack *webrtc.TrackLocalStaticRTP

	peerConnectionConfiguration = webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}
)

type WhipWhepRouter struct {
	log logg.Logger
}

func (r *WhipWhepRouter) RegisterRoutes(baseRouter *gin.RouterGroup) {
	baseRouter.Any(routeHandleWhip, r.handleWhip)
	baseRouter.Any(routeHandleWhep, r.handleWhep)

	var err error
	if videoTrack, err = webrtc.NewTrackLocalStaticRTP(webrtc.RTPCodecCapability{
		MimeType: webrtc.MimeTypeH264,
	}, "video", "pion"); err != nil {
		panic(err)
	}
}

func (r *WhipWhepRouter) handleWhip(ctx *gin.Context) {
	offer, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		panic(err)
	}

	mediaEngine := &webrtc.MediaEngine{}

	if err = mediaEngine.RegisterCodec(webrtc.RTPCodecParameters{
		RTPCodecCapability: webrtc.RTPCodecCapability{
			MimeType: webrtc.MimeTypeH264, ClockRate: 90000, Channels: 0, SDPFmtpLine: "", RTCPFeedback: nil,
		},
		PayloadType: 96,
	}, webrtc.RTPCodecTypeVideo); err != nil {
		panic(err)
	}

	interceptorRegistry := &interceptor.Registry{}

	intervalPliFactory, err := intervalpli.NewReceiverInterceptor()
	if err != nil {
		panic(err)
	}
	interceptorRegistry.Add(intervalPliFactory)

	if err = webrtc.RegisterDefaultInterceptors(mediaEngine, interceptorRegistry); err != nil {
		panic(err)
	}

	api := webrtc.NewAPI(webrtc.WithMediaEngine(mediaEngine), webrtc.WithInterceptorRegistry(interceptorRegistry))

	peerConnection, err := api.NewPeerConnection(peerConnectionConfiguration)
	if err != nil {
		panic(err)
	}

	if _, err = peerConnection.AddTransceiverFromKind(webrtc.RTPCodecTypeVideo); err != nil {
		panic(err)
	}

	peerConnection.OnTrack(func(track *webrtc.TrackRemote, receiver *webrtc.RTPReceiver) { //nolint: revive
		for {
			pkt, _, err := track.ReadRTP()
			if err != nil {
				panic(err)
			}

			if err = videoTrack.WriteRTP(pkt); err != nil {
				panic(err)
			}
		}
	})
	// !
	peerConnection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
		fmt.Printf("ICE Connection State has changed: %s\n", connectionState.String())

		if connectionState == webrtc.ICEConnectionStateFailed {
			_ = peerConnection.Close()
		}
	})

	if err := peerConnection.SetRemoteDescription(webrtc.SessionDescription{
		Type: webrtc.SDPTypeOffer, SDP: string(offer),
	}); err != nil {
		panic(err)
	}

	gatherComplete := webrtc.GatheringCompletePromise(peerConnection)

	// Create answer
	answer, err := peerConnection.CreateAnswer(nil)
	if err != nil {
		panic(err)
	} else if err = peerConnection.SetLocalDescription(answer); err != nil {
		panic(err)
	}

	<-gatherComplete

	ctx.Header("Location", "/whip")
	ctx.Status(http.StatusCreated)
	fmt.Fprintf(ctx.Writer, peerConnection.LocalDescription().SDP)
}

func (r *WhipWhepRouter) handleWhep(ctx *gin.Context) {
	offer, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		panic(err)
	}

	// Create a new RTCPeerConnection
	peerConnection, err := webrtc.NewPeerConnection(peerConnectionConfiguration)
	if err != nil {
		panic(err)
	}

	// Add Video Track that is being written to from WHIP Session
	rtpSender, err := peerConnection.AddTrack(videoTrack)
	if err != nil {
		panic(err)
	}

	go func() {
		rtcpBuf := make([]byte, 1500)
		for {
			if _, _, rtcpErr := rtpSender.Read(rtcpBuf); rtcpErr != nil {
				return
			}
		}
	}()

	peerConnection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
		fmt.Printf("ICE Connection State has changed: %s\n", connectionState.String())

		if connectionState == webrtc.ICEConnectionStateFailed {
			_ = peerConnection.Close()
		}
	})

	if err := peerConnection.SetRemoteDescription(webrtc.SessionDescription{
		Type: webrtc.SDPTypeOffer, SDP: string(offer),
	}); err != nil {
		panic(err)
	}

	gatherComplete := webrtc.GatheringCompletePromise(peerConnection)

	answer, err := peerConnection.CreateAnswer(nil)
	if err != nil {
		panic(err)
	} else if err = peerConnection.SetLocalDescription(answer); err != nil {
		panic(err)
	}

	<-gatherComplete

	ctx.Header("Location", "/whep")
	ctx.Status(http.StatusCreated)
	fmt.Fprintf(ctx.Writer, peerConnection.LocalDescription().SDP)
}
