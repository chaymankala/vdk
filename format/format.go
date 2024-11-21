package format

import (
	"github.com/chaymankala/vdk/av/avutil"
	"github.com/chaymankala/vdk/format/aac"
	"github.com/chaymankala/vdk/format/flv"
	"github.com/chaymankala/vdk/format/mp4"
	"github.com/chaymankala/vdk/format/rtmp"
	"github.com/chaymankala/vdk/format/rtsp"
	"github.com/chaymankala/vdk/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
