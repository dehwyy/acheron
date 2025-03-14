package routers

import (
	"errors"
	"net/http"

	"github.com/dehwyy/acheron/apps/stream_x/m3u8_broadcaster/internal/repos"
	"github.com/dehwyy/acheron/libraries/go/logg"
	"github.com/gin-gonic/gin"
)

const (
	contentTypeApplicationMpegURL = "application/vnd.apple.mpegurl"
	contentTypeMpegTs             = "video/MP2T" // ? Twitch uses `application/octet-stream` though

	streamName  = "streamName"
	segmentName = "segmentName"

	PlaylistRouterPath = "/:" + streamName

	routeGetPlaylistPath = "/playlist.m3u8"
	routeGetSegmentPath  = "/:" + segmentName
)

type PlaylistRouter struct {
	Log      logg.Logger
	FileRepo *repos.FileRepository
}

func (r *PlaylistRouter) RegisterRoutes(baseRouter *gin.RouterGroup) {
	router := baseRouter.Group(PlaylistRouterPath)

	router.GET(routeGetPlaylistPath, r.getM3u8Playlist)
	router.GET(routeGetSegmentPath, r.getSegment)
}

func (r *PlaylistRouter) getM3u8Playlist(ctx *gin.Context) {
	r.Log.Debug().Msgf("Request to playlist.m3u8 for <%s>", ctx.Param(streamName))

	playlistFiledata, err := r.FileRepo.ReadM3u8Playlist(ctx.Param(streamName))
	if err != nil {
		r.Log.Error().Msgf("Failed to read playlist: %v", err)

		statusCode := http.StatusInternalServerError
		if errors.Is(err, repos.ErrFileNotFound) {
			statusCode = http.StatusNotFound
		}

		ctx.String(statusCode, err.Error())
		return
	}

	ctx.Data(http.StatusOK, contentTypeApplicationMpegURL, playlistFiledata)
}

func (r *PlaylistRouter) getSegment(ctx *gin.Context) {
	r.Log.Info().Msgf("Request to segment for <%s> and <%s>", ctx.Param("streamName"), ctx.Param("segmentName"))

	segmentData, err := r.FileRepo.ReadSegment(ctx.Param(streamName), ctx.Param(segmentName))
	if err != nil {
		r.Log.Error().Msgf("Failed to read segment: %v", err)

		statusCode := http.StatusInternalServerError
		if errors.Is(err, repos.ErrFileNotFound) {
			statusCode = http.StatusNotFound
		}

		ctx.String(statusCode, err.Error())
		return
	}

	ctx.Data(http.StatusOK, contentTypeMpegTs, segmentData)
}
