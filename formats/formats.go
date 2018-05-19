package formats

var streamFormats = []string { 
	"hls",
	"dash", 
	"mss", 
	"hds",
}

var videoExtensions = []string { 
	"mp4", 
	"webm",
}

var videoEncodings = []string {
	"h264",
	"vp8",
	"vp9",
	"h265",
}

var playlistFormats = map[string]string {
	"hls" : "/hls/%s,%dp,%dp.urlset/master.m3u8",
	"hds" : "/hds/%s,%dp,%dp.urlset/manifest.f4m",
	"dash" : "/dash/%s,%dp,%dp.urlset/manifest.mpd",
	"mss" : "/mss/%s,%dp,%dp.urlset/manifest",
}

/* This box signals the bit rate information of the video stream. */
 type BtrtBox struct {
	Size               uint32			`datastore:"trackentry.btrtbox.size"`
	MaxBitrate         uint32 			`datastore:"trackentry.btrtbox.maxbitrate"`		/* the maximum rate in bits/second over any window of one second */
	AvgBitrate         uint32 			`datastore:"trackentry.btrtbox.avgbitrate"`		/* the average rate in bits/second over the entire presentation */
}


type TrackEntry struct {
	Width					uint16  		`datastore:"trackentry.width"`			// Track info (eg: 426)
	Height					uint16  		`datastore:"trackentry.height"`			// Track info (eg: 240)
	HorizontalResolution 	uint32  		`datastore:"trackentry.hresolution"`	// Track info (eg: 4718592)
	VerticalResolution   	uint32  		`datastore:"trackentry.vresolution"`	// Track info (eg: 4718592)
	FramesPerSecond      	uint16  		`datastore:"trackentry.fps"`			// Track info (eg: 1)
	BitDepth             	uint16  		`datastore:"trackentry.bitdepth"`		// Track info (eg: 24)
	Bitrate					BtrtBox
	FileName      			string				`datastore:"trackentry.filename"`
}
/*
	Desc    string    `datastore:"description"`
	Created time.Time `datastore:"created"`
	Done    bool      `datastore:"done"`
*/

type Video struct {
	Name			string		`datastore:"name"`
	Extension		string		`datastore:"extension"`
	Encoding		string		`datastore:"encoding"`
	Rate 			int32		`datastore:"speedrate"`		// Typically 0x00010000 (1.0)
	Volume        	int16   	`datastore:"volume"`		// Typically 0x0100 (Full Volume)
	Duration      	uint64  	`datastore:"duration"`		// Stream info
	Timescale     	uint32  	`datastore:"timescale"`		// Stream info (eg: for audio: 48000, for video: 60000)
	Language      	[3]byte 	`datastore:"language"`		// ISO-639-2/T 3 letters code (eg: []byte{ 'e', 'n', 'g' }
	HandlerType   	uint32  							// HDLR MP4 Box info (eg: 1986618469)
	Tracks        	[]TrackEntry
}

func N(){

}