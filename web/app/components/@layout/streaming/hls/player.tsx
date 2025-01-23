"use client"
import { API } from "@/lib/api"
import { VideoPlayer } from "@videojs-player/react"
import { useParams } from "next/navigation"
import { useMemo } from "react"
import "video.js/dist/video-js.css"

export function HLSVideoPlayer() {
    const { streamName } = useParams<{ streamName: string }>()
    const streamPath = useMemo(() => {
        return API.CreateHSLPlaylistURL(streamName)
    }, [streamName])
    return (
        <VideoPlayer
            id="video-player"
            src={streamPath}
            volume={0.6}
            preload="auto"
            liveui
            autoplay
            controls
        />
    )
}
