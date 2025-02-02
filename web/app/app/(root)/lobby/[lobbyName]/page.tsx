"use client"
import { RTCVideoPlayer } from "@/components/@layout/streaming/rtc/player"
import { StreamingLayout } from "@/components/@layout/StreamingLayout"

export default function Page() {
    return (
        <StreamingLayout>
            <RTCVideoPlayer streamName="dehwyy" />
        </StreamingLayout>
    )
}
