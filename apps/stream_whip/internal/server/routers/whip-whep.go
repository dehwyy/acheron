package routers

import (
	"io"
	"net/http"

	"github.com/dehwyy/mugen/apps/stream_whip/internal/rtc"
	"github.com/dehwyy/mugen/apps/stream_whip/internal/rtc/whipwhep"
	"github.com/dehwyy/mugen/libraries/go/logg"
	"github.com/gin-gonic/gin"
)

const (
	routeHandleWhip = "/whip"
	routeHandleWhep = "/whep"
)

type WhipWhepRouter struct {
	log logg.Logger
	api *rtc.API
}

func (r *WhipWhepRouter) RegisterRoutes(baseRouter *gin.RouterGroup) {
	baseRouter.Any(routeHandleWhip, r.handleWhip)
	baseRouter.Any(routeHandleWhep, r.handleWhep)
}

func (r *WhipWhepRouter) handleWhip(ctx *gin.Context) {
	offer, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		panic(err)
	}

	// TODO: extract from heades
	streamToken := "someToken"

	conn, err := r.api.NewPeerConnection()
	if err != nil {
		panic(err)
	}

	sdpAnswer, err := whipwhep.HandleWhipConn(conn, streamToken, string(offer))
	if err != nil {
		panic(err)
	}

	ctx.Header("Location", "/whip")
	ctx.String(http.StatusCreated, "%s", sdpAnswer)
}

func (r *WhipWhepRouter) handleWhep(ctx *gin.Context) {
	offer, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		panic(err)
	}

	// TODO
	streamToken := "someToken"

	conn, err := r.api.NewPeerConnection()
	if err != nil {
		panic(err)
	}

	sdpAnswer, err := whipwhep.HandleWhepConn(conn, streamToken, string(offer))
	if err != nil {
		panic(err)
	}

	ctx.Header("Location", "/whep")
	ctx.String(http.StatusAccepted, "%s", sdpAnswer)
}
