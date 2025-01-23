export class API {
    static CreateHSLPlaylistURL(streamName: string): string {
        return `http://localhost:8081/api/v1/${streamName}/playlist.m3u8`
    }

    static CreateWhepURL(): string {
        return "http://localhost:8082/api/v1/whep"
    }
}
