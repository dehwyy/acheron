package routers

import (
	"errors"
	"net/http"

	"github.com/dehwyy/mugen/apps/stream_broadcaster/internal/repos"
	"github.com/dehwyy/mugen/libraries/go/logg"
	"github.com/gin-gonic/gin"
)

const (
	contentTypeApplicationMpegUrl = "application/vnd.apple.mpegurl"
	contentTypeMpegTs             = "video/MP2T" // ? Twitch uses `application/octet-stream` though

	streamName  = "streamName"
	segmentName = "segmentName"

	PlaylistRouterPath = "/:" + streamName

	getPlaylistPath = "/playlist.m3u8"
	getSegmentPath  = "/:" + segmentName
)

type PlaylistRouter struct {
	Log      logg.Logger
	FileRepo *repos.FileRepository
}

func (r *PlaylistRouter) RegisterRoutes(baseRouter *gin.RouterGroup) {
	router := baseRouter.Group(PlaylistRouterPath)

	router.GET(getPlaylistPath, r.getM3u8Playlist)
	router.GET(getSegmentPath, r.getSegment)
}

func (r *PlaylistRouter) getM3u8Playlist(ctx *gin.Context) {
	r.Log.Debug().Msgf("Request to playlist.m3u8 for <%s>", ctx.Param(streamName))

	playlistFiledata, err := r.FileRepo.ReadM3u8Playlist(ctx.Param(streamName))
	if err != nil {
		r.Log.Error().Msgf("Failed to read playlist: %v", err)

		statusCode := http.StatusInternalServerError
		if errors.Is(err, repos.ErrorFileNotFound) {
			statusCode = http.StatusNotFound
		}

		ctx.String(statusCode, err.Error())
		return
	}

	ctx.Data(http.StatusOK, contentTypeApplicationMpegUrl, playlistFiledata)
}

func (r *PlaylistRouter) getSegment(ctx *gin.Context) {
	r.Log.Info().Msgf("Request to segment for <%s> and <%s>", ctx.Param("streamName"), ctx.Param("segmentName"))

	segmentData, err := r.FileRepo.ReadSegment(ctx.Param(streamName), ctx.Param(segmentName))
	if err != nil {
		r.Log.Error().Msgf("Failed to read segment: %v", err)

		statusCode := http.StatusInternalServerError
		if errors.Is(err, repos.ErrorFileNotFound) {
			statusCode = http.StatusNotFound
		}

		ctx.String(statusCode, err.Error())
		return
	}

	ctx.Data(http.StatusOK, contentTypeMpegTs, segmentData)
}
