package static

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type staticFilesFile struct {
	data  string
	mime  string
	mtime time.Time
	// size is the size before compression. If 0, it means the data is uncompressed
	size int
	// hash is a sha256 hash of the file contents. Used for the Etag, and useful for caching
	hash string
}

var staticFiles = map[string]*staticFilesFile{
	"steward_agent.swagger.json": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\xec[\xcdn\xe36\x10\xbe\xfb)\b\xb6\xc7\"\x0eҢ\x87ܜM\xbaH\xe1n\n'{(\x8a\xc0\xa0ű\u0085Dj\xf9\x93\xd4\r\xfc\xee\x85\xfe,J\xa2l\xc9\xcanW^\xf9b[⌾\xe1|\x1c\x0e9\xd4\xeb\x04!\xac^\x88\xef\x83ė\b_\x9c\x9d\xe3\x9f\xe2k\x8c\xaf\x05\xbeD\xf1}\x84\xb0f:\x80\xf8\xbe\xd2\xf0B$\x9d\x92\x88M\xb3\xdfK\xe2\x03\xd7g\x91\x14Z$\xb2\b\xe1g\x90\x8a\t\x1eKd?\x11\x17\x1a)\xd0x\x82\xd06y\x82'\xb82!(|\x89\xfeN\xa5H\x14\x05\xcc#\x9a\t>\xfd\xa4\x04\x8f\xdb>&m#)\xa8\xf1Z\xb6%\xfaI\x15Ч\t\xbc\xe2\x02B\xd8\am\xfdE\b\x8b\bd\xa2\xea\x96Ɛg4d|9gJ\xcfb\xd9̨\xa4\xa5\x04\x15\t\xae@\x95\x14 \x84/\xce\xcf+\x97\x10\xc2\x14\x94'Y\xa4\xb3\xae\x98!e<\x0f\x94Z\x9b\x00\xe5\x9a\xce,\xf5\x89\x90\xf2\x9e $5e\b\xe1\x1f%\xacc=?L)\xac\x19g\xb1^\x95\xbba\aw\x91)\xc6%\xf1\xad\xf5ok?\x11SX\x13\x13\xe8\xc3\xe892\x1c\xfe\x89\xc0\xd3@\x11H)\xe4Έ\xbe6H\xc35\v\xe1&V\xba\a\xf7\xc4a\x01\x8e\x88$!h\x90\x05;\xd2O\xc5\x1cN\u008c\xc2D\xea*`\x96\x98\xf8ـ\xdcToI\xf8l\x98\x84\x98\x19k\x12(\xa8\xdc֛(S+\x19\xf7\xab\xc2k!C\x12\xf7-f\\\xff\xfa\vn\xf2B\x03ֈ\xf8\xb0T\xec_\x18\b\xde\xe4\xfbKAu2\xe1\xd1b\x82&~\x95\x03\xe9X.D\x1f'\x15cp$T\x9bh\xf0N\x02\xd10\xa0x`\x01\xfe\xde\"B\x17\x1eL,\xe1|\xae\x98\xbe\xa6S\x1a\xa3[{\xdahɔ\x8f\x11\x1d\x16S,\xc0\xe3ܱ?\xbe\xe5\xbcpǸ8\xf3\xd8\x13\xe2\xb44_3\xc2\xeda\xf6\x92\xd1\xed\x94\xf1g\xa6I桎\xb9\xd1{з;\xf1߄\x1c\x10\xdb\x13\xa8\xb7\x9c\xe5\xe8Gʷ\xa0\xfc\xf24(/\xc1gJ'+\x9d\xaeQ}\x91\x89\xfeiV\x01\xf3\xaeo\xaf\x87\xc1\xf6\x1dܑ\xe6\xa7N\xf3J\xb6\xd2:\x90\x0f(v\xe7hG2\xef'\xf30h\xbc[\x82Q\b@C\v\xb6^'\r\aDX\v\xf0\xc8\xd9S\xe0lc\xe8\x9d\x06\xc4p\xef\xe9\x88\xccb\x9e\b\x0ei\xa7\xb1\x00<\x92\xfa\xb4I\xad\x9e\x8c\xa6\xe2\x85\x1fA\xeb\xfbLt@\xc4.A\x1e\xa9}zԮ\xf5g\x87R\xd0}*;\x98ZP\x8awd\xf1X\f:\xf9bP\v,yA\xfa\x84kSC\nP6\xe2\xb1:\xd5m\xfa:~\xa7gH\f\xd9\xc1\x1dg\xb0\xefy\xafgH\x94\xb5\x11\x8f\xac=\xd5\xd5\xc3\xf45\xfd\xeeu:`H\xb4\xb6\x11\x8f\xb4>\xb0\x9cȩ1\x1cv\xef\xcecZ}V\x9c\x9eLNu\xae\xccz\xc676\xd9sXb\xf5\t\xbcb\xfd\x147\x8f@jV\xa1q\xd2~idP%w\x93y\xb6\x0f\x9eI`\xe0\x80`\x89tŲe\xb5\xd1\x16U\xb7\xcea]\xa2G\x0f\x13\xa1\xa2\xa0\xb5}\x9e\xa0\x8d\xe61\xae\xc1\a\xd9d\x1f\xe3\xfa\xe7\v\xb7\xd6\x10\x94\">\x1c\x03\x88\x82&,PM\xa2DJR^-a\xa6!T\xf5A\xee\x1e\x8f6\xa5\xdc\x03\xd0\xe9'\xfb\xf8F\x1f?1zL\x9fd\x03\xbc\xb3\x1cQ\x8a\xf9\x1c\xe82\x8d\f\xcb\xe3\x9e\x0e\x9c\n\xa9\xc8*\x80B\xcf\x1b\xf8\xa7Mp9\xec\f\xd7Y\x9a\x1e\xfeY\t\xbai\xddG{\xc1i-\xd9\xca\xe8^h\x8e\xf5{֦$\xb7\xffXR\x8e\xf6!\x96\xech\xe1C\xf9iM\xd1\x11\x037aiv\xc0\xf7\x0f\x8b\xdb\x0f\xef\xedD\xe4\xc3\xc7?\xaen\x16\xf6\x95\xbb\xab\xdfo\xde=\xd8Wf\x8b\xc5\xec/\xfb\xc2\xd5\xddݼ\xacd>ύx\xb4r\xff<\x1bȟ\xeb\xb2\xcaup\xf5\xeb\x8c\xf7}}\xedܱ\xf8\xffa\xb9\n\xccͨ\x9a5\xb46̥\xa2v,\xa3G\xbf\x90J|oy\xa0\xafew\xd5w\x15z@u&\xa2\a*L\xa9H;\xb0\xae:k7\xc7\xd4\xdf\t\xe9a\xae'L\xdd3\xed\xf2\xb0\xca\xf6qi\x82\xac\xbe\x9d\U000e6e46\x9b\"\xad\xe77G\x1d\xe5\x9b\xeb@7\v߸\x03\xab\xbcm݃\xf5\xf3\x87\xdd\x18|_\xab\xd8}\xdb\xf9^\xf1\xe2[\xf7T1\x9fͿ\xf8h\xd8%F\xdd\xfd\xe9.\x91w\xf3\xa9\xeb}\x83c4t\x9c\xb3\xe2\x85\xeed;\xf9/\x00\x00\xff\xff\x99\x11d-\xf28\x00\x00",
		hash:  "871990a7c36b7c34dca63421326134cb7b51e07b3e28f315e4e6a5ec8048f66a",
		mime:  "application/json",
		mtime: time.Unix(1592403827, 0),
		size:  14578,
	},
}

// NotFound is called when no asset is found.
// It defaults to http.NotFound but can be overwritten
var NotFound = http.NotFound

// ServeHTTP serves a request, attempting to reply with an embedded file.
func ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	path := strings.TrimPrefix(req.URL.Path, "/")
	f, ok := staticFiles[path]
	if !ok {
		if path != "" && !strings.HasSuffix(path, "/") {
			NotFound(rw, req)
			return
		}
		f, ok = staticFiles[path+"index.html"]
		if !ok {
			NotFound(rw, req)
			return
		}
	}
	header := rw.Header()
	if f.hash != "" {
		if hash := req.Header.Get("If-None-Match"); hash == f.hash {
			rw.WriteHeader(http.StatusNotModified)
			return
		}
		header.Set("ETag", f.hash)
	}
	if !f.mtime.IsZero() {
		if t, err := time.Parse(http.TimeFormat, req.Header.Get("If-Modified-Since")); err == nil && f.mtime.Before(t.Add(1*time.Second)) {
			rw.WriteHeader(http.StatusNotModified)
			return
		}
		header.Set("Last-Modified", f.mtime.UTC().Format(http.TimeFormat))
	}
	header.Set("Content-Type", f.mime)

	// Check if the asset is compressed in the binary
	if f.size == 0 {
		header.Set("Content-Length", strconv.Itoa(len(f.data)))
		io.WriteString(rw, f.data)
	} else {
		if header.Get("Content-Encoding") == "" && strings.Contains(req.Header.Get("Accept-Encoding"), "gzip") {
			header.Set("Content-Encoding", "gzip")
			header.Set("Content-Length", strconv.Itoa(len(f.data)))
			io.WriteString(rw, f.data)
		} else {
			header.Set("Content-Length", strconv.Itoa(f.size))
			reader, _ := gzip.NewReader(strings.NewReader(f.data))
			io.Copy(rw, reader)
			reader.Close()
		}
	}
}

// Server is simply ServeHTTP but wrapped in http.HandlerFunc so it can be passed into net/http functions directly.
var Server http.Handler = http.HandlerFunc(ServeHTTP)

// Open allows you to read an embedded file directly. It will return a decompressing Reader if the file is embedded in compressed format.
// You should close the Reader after you're done with it.
func Open(name string) (io.ReadCloser, error) {
	f, ok := staticFiles[name]
	if !ok {
		return nil, fmt.Errorf("Asset %s not found", name)
	}

	if f.size == 0 {
		return ioutil.NopCloser(strings.NewReader(f.data)), nil
	}
	return gzip.NewReader(strings.NewReader(f.data))
}

// ModTime returns the modification time of the original file.
// Useful for caching purposes
// Returns zero time if the file is not in the bundle
func ModTime(file string) (t time.Time) {
	if f, ok := staticFiles[file]; ok {
		t = f.mtime
	}
	return
}

// Hash returns the hex-encoded SHA256 hash of the original file
// Used for the Etag, and useful for caching
// Returns an empty string if the file is not in the bundle
func Hash(file string) (s string) {
	if f, ok := staticFiles[file]; ok {
		s = f.hash
	}
	return
}
