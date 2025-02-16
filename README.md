# Mugen

## Streaming
There're 2 types of streaming:
1. ***SRT*** - at least 2 seconds latency
2. ***WebRTC(Whip)*** - very low-latency as connection established P2P

- Tested on my PC.
- No lags.
- No freezes.
- High quality (there are limitations, I guess).
- ***120***  FPS.
- (SRT) ***Use FFmpeg AAC as audio encoder!*** . Otherwise audio won't work! *P.S.* on **2nd** screenshot other codec is set.
- (Whip) ***Use FFmpeg Opus audio encoder.***

### OBS SRT Stream Settings
<img src="docs/screenshots/obs_stream_settings.jpg">

### OBS SRT Output Settings
<img src="docs/screenshots/obs_output_settings.jpg">

## OBS Whip Stream Settings
<img src="docs/screenshots/obs_whip_stream_settings.jpg">

### OBS Whip Output Settings
<img src="docs/screenshots/obs_whip_output_settings.jpg">


### StreamHost service
- Responisible for handling data passed by SRT connection from OBS.
- Recommended `SegmentDuration`=1000 (ms) and `SegmentsPerStream`>=10 (via `config/config.toml`)

## App screenshots

### Auth page
<img src="docs/screenshots/app_login.jpg">

### User page (not dedicated for watching streams)
<img src="docs/screenshots/app_stream.jpg">
