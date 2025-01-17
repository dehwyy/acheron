import { VideoPlayer } from "@videojs-player/react"
import { useMemo } from "react"
import "video.js/dist/video-js.css"

interface Props {
    streamName: string
}

export function StreamingLayout({ streamName }: Props) {
    const streamPath = useMemo(() => {
        return `http://localhost:8081/${streamName}/playlist.m3u8`
    }, [streamName])

    return (
        <VideoPlayer

            id="video-player"
            src={streamPath}
            liveTracker={{
                liveTolerance: 3,
            }}

            volume={0.6}
            preload="auto"
            autoplay
            controls
        />
    )
}
