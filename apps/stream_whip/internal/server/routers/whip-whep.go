package routers

import (
	"net/http"

	"github.com/dehwyy/mugen/apps/stream_whip/internal/rtc"
	"github.com/dehwyy/mugen/apps/stream_whip/internal/rtc/whipwhep"
	"github.com/dehwyy/mugen/apps/stream_whip/internal/server/extractors"
	"github.com/dehwyy/mugen/libraries/go/logg"
	"github.com/gin-gonic/gin"
	"github.com/pion/webrtc/v4"
)

const (
	routeHandleWhip = "/whip"
	routeHandleWhep = "/whep"

	ctxOffer = "offer"
	ctxToken = "token"
	ctxConn  = "conn"

	headerXStreamName = "X-Stream-Name"
)

type WhipWhepRouter struct {
	log logg.Logger
	api *rtc.API
}

func (r *WhipWhepRouter) RegisterRoutes(baseRouter *gin.RouterGroup) {
	router := baseRouter.Group("/")
	router.Use(r.prepareConnection())

	router.POST(routeHandleWhip, r.handleWhip)
	router.POST(routeHandleWhep, r.handleWhep)
}

// Middleware
func (r *WhipWhepRouter) prepareConnection() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		offer, err := extractors.BodyToString(ctx.Request.Body)
		if err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		// TODO: make authorization
		token, err := extractors.GetAuthorizationToken(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		conn, err := r.api.NewPeerConnection()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.Set(ctxOffer, offer)
		ctx.Set(ctxToken, token)
		ctx.Set(ctxConn, conn)

		ctx.Next()
	}
}

func (*WhipWhepRouter) handleWhip(ctx *gin.Context) {
	offer := ctx.GetString(ctxOffer)
	token := ctx.GetString(ctxToken)
	conn := ctx.MustGet(ctxConn).(*webrtc.PeerConnection) // nolint: revive

	// TODO: get `StreamName` from auth
	streamName := token

	sdpAnswer, err := whipwhep.HandleWhipConn(conn, streamName, offer)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"desc":  "Failed to handle WHIP connection!",
		})
		return
	}

	ctx.Header("Location", "/whip")
	ctx.String(http.StatusCreated, "%s", sdpAnswer)
}

func (*WhipWhepRouter) handleWhep(ctx *gin.Context) {
	offer := ctx.GetString(ctxOffer)
	// token := ctx.GetString(ctxToken) // TODO: somehow validate user session
	conn := ctx.MustGet(ctxConn).(*webrtc.PeerConnection) // nolint: revive

	streamName := ctx.GetHeader(headerXStreamName)
	sdpAnswer, err := whipwhep.HandleWhepConn(conn, streamName, offer)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Header("Location", "/whep")
	ctx.String(http.StatusCreated, "%s", sdpAnswer)
}
