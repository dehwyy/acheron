"use client"

import { API } from "@/lib/api"
import { useParams } from "next/navigation"
import { useEffect } from "react"
import "video.js/dist/video-js.css"

const authToken = "dehwyy"

export function RTCVideoPlayer() {
    const { streamName } = useParams<{ streamName: string }>()

    useEffect(() => {
        const conn = new RTCPeerConnection()

        conn.addTransceiver("video", { direction: "recvonly" })
        conn.addTransceiver("audio", { direction: "recvonly" })

        conn.ontrack = (event) => {
            const video = document.getElementById("video-player") as HTMLVideoElement
            video.srcObject = event.streams[0]
        }

        conn.createOffer().then((offer) => {
            conn.setLocalDescription(offer)

            fetch(API.CreateWhepURL(), {
                method: "POST",
                body: offer.sdp,
                headers: {
                    "Authorization": `Bearer ${authToken}`,
                    "Content-Type": "application/sdp",
                    "X-Stream-Name": streamName
                }
            }).then(r => r.text()).then((answer) => {
                conn.setRemoteDescription({
                    sdp: answer,
                    type: "answer"
                })
            })
        })
    }, [])

    return (
        <video
            className="w-full h-full"
            id="video-player"
            autoPlay
            controls
        />
    )
}
