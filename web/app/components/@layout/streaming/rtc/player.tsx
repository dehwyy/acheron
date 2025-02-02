"use client"

import { API } from "@/lib/api"
import { useEffect } from "react"

const authToken = "dehwyy"

interface Props {
    streamName: string
}

export function RTCVideoPlayer({ streamName }: Props) {
    // const { lobbyName } = useParams<{ lobbyName: string }>()

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
            className="h-full w-full"
            id="video-player"
            autoPlay
            controls
        />
    )
}
