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
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\xec\\oo\xdb8\xd2\u007f\x9fO1\x8f\x9e\x03\x9a\xdc\xfa\xe44\xbb\xd8\x17Y\x04\xb8l\xffm\x16\xd96H\x93\x17\x87*H(idqK\x91*I9\xf5\x16\xf9\xee\x87!%[\xb6e\xc7N\xdan\xdds\x81±H\x0e\u007f3\xf3\x9bᐒ\xfci\a 0\xb7l0@\x1d\x1cBp\x10\xee\a=\xba\xc6e\xa6\x82C\xa0v\x80\xc0r+\x90ڟ1\xc9\r\x1c\xa7\x05\x97p|v\xe2\xfa\x02\x04CԆ+I=\xf6\xc3\xfd\xf0i\xb0\x03p\xe7\xe4$J\x9a\xaa@\x13\x1c\xc2;ߗ\x95\xa5\xe0\t\xb3\\\xc9\xfe\x9fFI\xea{\xe5\xfa\x96Z\xa5U\xb2b_fs3\x01\xd8g\x03\x94vr\x01 \x18\xa0m}\x05\bT\x89ډ:I\t\xa8S\xe2\xfa\x94\x1b{LckU\\O\x8d\xa6TҠ\x99\x12\x00\x10\x1c\xec\xef\xcf\\\x02\bR4\x89業\rp\f\xa6J\x124&\xab\x044\x92\u0096x7\xc8$9\x16lN\x18@\xf0\x0f\x8d\x19\xc9\xf9\xff~\x8a\x19\x97\x9c䚾\xb1x\xcbt:\x86{^\v\x0e\xa6\x86ߵ\xbeݵg\fR\xccX%\xec\xfd\xe8%T\x12?\x96\x98XL\x01\xb5Vz\xac\xc4cuЕ\xb4\xbc\xc0\x17$t\t\xee\x9d\x0e\r\x82\x92iV\xa0E=a\x87\xff7\xa3\x8ed\x85#\xaa\xb1L\xdbY\xc0ܩ\xf8\xa1B=\x9am\xd2\xf8\xa1\xe2\x1a\x89\x19\x19\x13\x06g\x9a\xed\xa8\xac\xc5j.\a\xb3\x833\xa5\vF\xb6\r\xb8\xb4?\xff\x14,\xf2\xc2\x02\xac%\x1b\xe0\xb5\xe1\u007f\xe1\x86\xe0u\x9f_\nj'\x13\xaeZL\xb0l0\xcb\x01\x1f˓\xa1W;3\xca\x04\xa52\xabd\x83g\x1a\x99\xc5\r\xca\a-\xc0ی\xb0\x9c\xb5\xb1JGݬ\xedji\x91\xd6\xea\n?\x8f\xb7<\xb1V\xd0\xf3\xa1|\xdfi\x19\xa9Y\x13\xfb\x9f\xdcg\xc8ӻ\xf6\xf2\xb8bD\\\x96\xe9fED\v\xf06\"\x96GDË\uea20\nk\xbd\xa8X\x92\xc9{\xdb\xf8\xbc'>\xafyz\xd7\xe7r\xc8-\xaby\xb6f%\xfb\n\xed\xc9x\xf8K\xa57(f\x1d\xd4\x13\xc9\x1b\xf4\xdb\xc0]!p\xaf\xbfJ\xe0~q\xcak\x1cpc\xdd\xeesݵ\xe9\xbc\x1ezVł'\xcfO\x9eo\x06\xdb\xc7p\xb74\xfffh\xbe\x19\xebS\x8b9\x1f*4\u007f\xc3R5SD\xae\xbc2m\xd0bԠ\xddF\xe7\xf2x،\xe5g|\x02\x90\xa2@\x8b+\xb0\xf5\xb9\xeb\xb8A\x84m\x01\xder\xf6{\xe0\xec\xc2\xd4\xdb\x17\xac\x92I\xfe\x80R\xe9\xd4\rܤ\x83\xee\t\xe0-\xa9\xbfoR\x9b\xbc\xb2\xa9\xba\x95\x0f\xa0\xf5\xdbz\xe8\x06\x11{\n\xf2\x96\xda\xdf\x1f\xb5\xe7\xec\xb9Ɲȷ~\xec\xc6܊\xf4x\xb7,\xdeދ\xfc\xee\xefE\xae\x80\xa5y\n\xe2;\xbe5\xbaI\t\xaa\x8dx\x9b\xa2\xbe\xf9í\x9aZ_\xedH\xcb\xc3|\xf8\x91\xd6&\x85\xc2\x18\xee6\x0e\xfe\x97\x0f\xb56\x89\xb2m\xc4[\xd6~\xafۤ\xfe'\xff\xf9\xa8\xa7S6\x89\xd6m\xc4[Z߳oj\xa8\xb1\xbd\x01\xf8\x95k\xa4\xf1\xe3\xdb-\f\x93\x87\xadK\xad\xac\x8a\xab\xecX\x8e\xda1\xdbXW\xc5\u007fb2\xd9\xefR\xf7\x12\xb5\xe53\xd1\xe8\xfa_WZ\xcc\xc6\xe8\x92\xfd\xe2|\xa8^\x9e\x9f\xf65\x1aU\xe9\x04\x81<\x056g\x16*\xc9?T(F\xc0S\x94\x96g\x1c\r\xd8\x1c\x81d\x83\xca\xdc\xdf\x065g\x82\xff\x85i$\x9dJ\x89\x12\x10WY\x86\x1a\n4\x86\r0\x84\x8b\x1cA0c\xc1\xe0\xa0@i\x9b\xc1\x97\xe7\xa7O\f\x10\x05\xa1\xa8\x8c\x05\x8d\xa5F\x83\xd2F\x92\x9a\xb3J\x88\x11|\xa8\x98\xa0\xb9S\x8f\xac\x1e\xea0\xec2\x03\\F\xf2\x86D\xf4\aJ\r\x04\x86\x8da\xc3\xe7\x95\xcfu7{\x1e\x81\x1bnrU\x89\x14b\x04.\x81A¤\x92<a\x02h\v\x1d\xc9]\f\aa\x0f\x04\xb2\x94\xcb\x01DA\x18\x05\xc0\rHe\x81%\t\x96\x16ӽ0\x92\x91<\x91Pj\x96X\x9e`\x0f,\xb2\xc2@e*F\x88K\x8d\x89*J.h\x12\xab\x1cޘK\xa6G\xc0\x84pЍ\xb7\xb0\xcdq\x14\xd5Y\x06\xb8\x05\xab\xa02\x0e\x1a\x8dI\x94\xb4\xf8\xd1Y\xebX\x8eB\xf8M\xdd\xe2\x10u\x8f\xb0\x92\xed\f\xdc\xe6<\xc9\xdd\x10\x9bc$]@ \xdc\xe4֖7=\xffinz\xa04H\x05\xbe\xb5\aJ\"\xe9\r\xca1\xc0!6h\xa1*\x819l\x914\xa8\x87\xa8=Ă\x95\xc6[\xdb\xcdhU\xe3Vh\xd1\x1a\x98\x81L\t\xa1n\xcd!\x19\xe7\x9fp\x92M\xa6$\x03\x96Z\ry\x8a\xe9\x18\x15]d\xc6T\x05\xa6!\r8\x96\xf0\xdb\xc5\xc5\x19\xbczq\x01J6\xf4\xf0\xbc\x18q\x14)0x7\xeb\xe2\x8bQ\x89W\xef\xae\"\t0d\xa2r\x96\xab-\xed\x8fD\x9c\xee\xf5+\x11\xc0\xa4\xcf\xe3~\xbe\xc9[\x11\x06\x98Fr\x8d\xbaŔ4LXB\x8cU\xea}URү\x845\x103\x83i\r\x8d&\xbc<?u\xd2s6t\xe6/Z~O\xbd\xe3Y\x03\x86\xfe\x1e*\x9e\x02\x93#\x1a\xebE;Zj̔\xc6^ӓ\x040\xcbc.\xb8\x1d\x81DL\x9d\xcdc\x04\x17\x1azH\x81\x06\x04#ə\x1c\xa0ku\x8c\na\xf7\xd2 \xd4\xe7'\xa4\b9\x8dH\xef\xbd\xc6$\x1b8\xe0\xb1F\xf6\x9e\xd8]K\b\xf7\xc8e\xaf\x95\xc5C\xb097\x90U2\xf1\xd4 \f5\xfb\x93Jk\x94V\x8c\x80\r\x19\x17,\x16c\x9e\xaa,\xe3\tg\xa2\xce\x00q\x95\x81F\x81\xcc`\x0f\x98L\x89ص\x90\x8aLH\xec\x9d\x10*\xc6\x01\x97\x92\xe0\xdcr\x9bG\x92ZB\xefgVr\x13&\xaap\xf1\xe6\xb26\x1aP6\xf7Ԕ\xb3<\x87]\x12\x9c#`Q\xdaQͽ=(\xf8 \xb7\x10c$\xdd\xec4\v\xf0\xa2\x14H\x99\xc8\xf9\x1fL\x89\t\xcfx\x02\x06\v&-OL\x18t.\xbe\x8ed\xeb\xa4\xda\xc9\xc1\\<\xb2\xb84\r\xffAL\x8f\x11\x18Q\x99\xa7\xad\xcc\n\xb3\x89\xb5\u0381,VCl\xc0\xd7\x0eo\x03\xef\xd8\xe9L\xcdxs,G7.\xcb0NA \x81\xe9\x98[M<\\2{\x13\xffL\xa8\xdak\xc0\"I\xc1\xea\x12\x86\x9f$^\xbaV\x8c\x97\x06\xf2\xecYC\x1a\xc1c7w\x9d+\f\x98\xaa,\x95vi\xb1d\xc9\xfb~%郒\xa1\x0fw\xd30\xd0'o\x95Ae}\xe04\x146\x14\xa1,M\xb9\xe73\fPR\x05\xec\x10\xd8\\\xa5\xa6\xc1F2\x9d\xfd\bы\x8f\x8c\b\x02O\x0f\xe1\x8c&$\x12\xd7s\xb3\xb1\xfa\\³\x1f~p\xfdɸ/\x95\x82L)8\x820\f\u007f\xf1\xd7H(\x93\xa3\xfa\x1b\x93\xa3\x90ĽԪ\xd8͔ګ\xaf\x87a\xe8\xff\xe0\x19\xecR\xa7K7Յڍ\xaa\xfd\xfd\x83\x9f\xa9\xeb\x1e|\xf2}Z\xdd\xef\xdaP\x0f\xee\x81\xfa;\x1b\xb2U\xb0\u0091[kH\xc0R\x8c\xdc\xec\xbeT*L\x043\xa6\x8d\u038b\xa5\x1e\x1eE\xab\xd7/-\xd8\xd0\xe0\xfe\xf1\x1e\xdcg#\x9b+9F\xeeſTj7\f)o\xd5v\xf5\xa8w\xf7\xa6\r\xed\x14\x98\xc7O\xcd'\x1e\xfe\xf3\x17o\x9f\x9d\x9f\x9c]\xbc9\xdf;l4\x98x\xa05\xbe\x96\xd0\x02\xfe\xd3=\xc0_\xa9\x06\xb3\x03}x\x04ޛe\x1c\xbeT\xeaS\x18\x86wu3\x93\xa3\x1e-Lԧ\xf4\xa9\xfc\x0f\xa6M\xce\x04\xe9\xd4\xc20V\xa2Sb#\x8eg3\xc2.e1\x11\xe7&s\x8eu\xbd\xfe\xef\b$\x17\x13\xf7\xb5\xe6p~\xa2\xba\xc9\xe9քK\xb3\x8eC<\x82r6po\xb9\x10\xd4P\xef\xb1(\xddG\xf2IGF\xefSi\x17\xba\x06Z\xa0\x9eP\xfd0\xce\x16\x94IȪt\xc1[6\x92\xe3h\x95b\xd4\xd4;s\xf5\xe1x\xc1\x03\x96Y\xf4k\x81+;\x9f\xf4\x9fD\xb2N\x15\xcd\x14\xbe\x8a\xc2ڛQ\x90)\x15\xc6L;t\x1f\xfb\xa3\xf0\xaf(\xf0\xfa\xf8\xe2\xc3\x17FNx\x14\xb8VG\x87H\xfe\xfe\xf6\xcd\xebH\x1e\x1d\x1d\x1dyk\xd1\xf7I!\xeb\xd7\x17E\xa4\x03\x9fn}\x9dR\x99:?j\x1cT\x82\xe9H\xce\x0f\xa1\xe6\x14'I\xb3\aXĘ\xa6\x93\xf4٫\xb3\xaf\x8cd+\xc7e\x0e\xf0Ϳ\t\xf2M]\"\x8e\x93|\xdb\x04aC\xe6Æ\xaadl\xe2\xef\xa4\xceʸ\xc0:p\x1br\x9f\xa16JN8\xe3\xd7=ȸ6\xf6\xdaY\xe8\b\x9e\xfe2\xd3J~h\x1a\x0f\xa62\x01\xed\x15\x9b\xceQ\xe0PG\xc1!DA\x17o\xa6\x81\x85\x1eJ\x14\xf4&\x02\x1c\x8c\u05ec\xf0B\xaa\xfd\xfd\x1f\x13\x0f\xc1\xfd\x8d\xad\x9e\x04iq\xc7\x16ē\xac.+\xa6\xad\xef\xed\xc8\rܢ\x10\xffz/խt\xbc͙\xa1\x9dEe\xac*\xc0\xd3cڹ=\xbfP\xcex\xdc\aOk\x1ar\xa9\x1c\x00\xf3\x0e\x8d䍣N\xe3\xd1\\\x89Ի\xb35\x93\xdb\xd4\xd4L\x80z\xe7Q\x13!\x92N\xcc\xd8\xe7\xb0K\xfcoTy\xb7h\xf3t\xf5\xeej\xef\xf01~\x9a\x167\xe5*\xa7\x8f\x97\xf14<xz`\xa2\xa0\xb6z0u\xcc5u\\\xf2\x88\xbd2\xce\bXz\x9c\xd1.\xfb\x12\x95.\xac\xfa\xb8\xb48@\xbd\xa8\xec\xe3\xd2\xfex\xd0-\xb5\xb6\xfdC\x00\xa5h\x19\x17f\xd1P\xa65\x9b>K\t\xb8\xc5\xc2\xcc\x1fzu\x1f\x94\xb4\xcf&\xba\x0f\xa4:\x8f#\xa7^\x00z\x84\x9fx\xfa\x10\x9b\xd4'Lk\x8fc\xc6\xf0\x81\xc4\xf4\xda\x1f']?lv\x94\xa9҆vE\x139\x9f\xc1?\xab\x1c%\xdf\uf32e\xb7\x82\x1e\xe1\x1fwR\xb7\xaa\x8d\x96\x82\xb3V\U000f8c8fB\xf3P\xbf\xd7}\xa6\xc6-\u007f\xc1\xaaA{A#\xd7\xd4\xf0bz\xb6E\x9b\xc6\x00eUL\x1d3\x06o/\xceO^\xbfj\x1f̿\xbe\xfc\xe3\xd7\x17\xe7\xed+o~\xfd\xfdų\x8b\xf6\x95\xe3\xf3\xf3\xe3\xff\xb4/\xfc\xfa\xe6\xcd鴐\xd3\xd3F\x89\xab\xd6\x0e\xb19\x1do\xe6\xedҪ\xeb\x85\xe9\xaf\x13\xef\xcbl\xdd\xf9\xa8\xc2\xdf\x0f\xab\xeb\xc9\xf2Ũ\x16KXY\xb1.\x11s\xefc<\xc2.l&\xbf\xaf\xfdJ\xe8Rs\xcd\xdfe\u007f\x04\xd4\xce;\x04\xeb\xdd\x1dX\n\xb6\xeb\x01\xeb\xf5\x1c3\xff[$\x8fP7QռgV;\x9e\x9aynlj\x81\x9c\xfdU\x98\xcfZk,zkx\xc5\xf5\xad\xe3\x01\xcao\u0380\xdd,\xfc\xcc\x06\x9c\xbf\xab\xb5\xa2\x05\xe7އ{ljx`\x05\x95>l\xd8\x10\xf5{\xfc<\xe5\xc8\xfc;\xa5\xeb\xc5\xf2۹\x87\x96\xbf\xed\xcaw\xf2\x83S\xeb\x17\xcdM]\xf3\xc5\xf3¸D\\\x9f\xd9\xddo\t\xac\xe7Ӯ_\xc2x\x88\x845W\xef\x1d\xfa\u007f\xb7\xf3\xdf\x00\x00\x00\xff\xff\x89`\xea\xd9ZL\x00\x00",
		hash:  "a310eb812b9a06862d047b66ae9d39033ec21514f44ab05f05d7b555d79e9f30",
		mime:  "application/json",
		mtime: time.Unix(1592500542, 0),
		size:  19546,
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
