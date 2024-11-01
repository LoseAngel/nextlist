package utils

import "strings"

var mimetypes map[string]string = map[string]string{}

func FindMimetypeByExt(ext string) string {
	mimetype, ok := mimetypes[strings.ToLower(ext)]
	if ok {
		return mimetype
	}
	return ""
}

func newMIME(mimeType string, ext string) {
	mimetypes[ext] = mimeType
}

func init() {
	newMIME("application/x-xz", ".xz")
	newMIME("application/gzip", ".gz")
	newMIME("application/x-7z-compressed", ".7z")
	newMIME("application/zip", ".zip")
	newMIME("application/x-tar", ".tar")
	newMIME("application/x-xar", ".xar")
	newMIME("application/x-bzip2", ".bz2")
	newMIME("application/pdf", ".pdf")
	newMIME("application/vnd.fdf", ".fdf")
	newMIME("application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", ".xlsx")
	newMIME("application/vnd.openxmlformats-officedocument.wordprocessingml.document", ".docx")
	newMIME("application/vnd.openxmlformats-officedocument.presentationml.presentation", ".pptx")
	newMIME("application/epub+zip", ".epub")
	newMIME("application/jar", ".jar")
	newMIME("application/x-ms-installer", ".msi")
	newMIME("application/octet-stream", ".aaf")
	newMIME("application/msword", ".doc")
	newMIME("application/vnd.ms-powerpoint", ".ppt")
	newMIME("application/vnd.ms-publisher", ".pub")
	newMIME("application/vnd.ms-excel", ".xls")
	newMIME("application/vnd.ms-outlook", ".msg")
	newMIME("application/postscript", ".ps")
	newMIME("application/fits", ".fits")
	newMIME("application/ogg", ".ogg")
	newMIME("audio/ogg", ".oga")
	newMIME("video/ogg", ".ogv")
	newMIME("text/plain", ".txt")
	newMIME("text/xml", ".xml")
	newMIME("application/json", ".json")
	newMIME("application/json", ".har")
	newMIME("text/csv", ".csv")
	newMIME("text/tab-separated-values", ".tsv")
	newMIME("application/geo+json", ".geojson")
	newMIME("application/x-ndjson", ".ndjson")
	newMIME("text/html", ".html")
	newMIME("text/x-php", ".php")
	newMIME("text/rtf", ".rtf")
	newMIME("application/javascript", ".js")
	newMIME("text/x-lua", ".lua")
	newMIME("text/x-perl", ".pl")
	newMIME("application/x-python", ".py")
	newMIME("text/x-tcl", ".tcl")
	newMIME("text/vcard", ".vcf")
	newMIME("text/calendar", ".ics")
	newMIME("image/svg+xml", ".svg")
	newMIME("application/rss+xml", ".rss")
	newMIME("application/owl+xml", ".owl")
	newMIME("application/atom+xml", ".atom")
	newMIME("model/x3d+xml", ".x3d")
	newMIME("application/vnd.google-earth.kml+xml", ".kml")
	newMIME("application/x-xliff+xml", ".xlf")
	newMIME("model/vnd.collada+xml", ".dae")
	newMIME("application/gml+xml", ".gml")
	newMIME("application/gpx+xml", ".gpx")
	newMIME("application/vnd.garmin.tcx+xml", ".tcx")
	newMIME("application/x-amf", ".amf")
	newMIME("application/vnd.ms-package.3dmanufacturing-3dmodel+xml", ".3mf")
	newMIME("image/png", ".png")
	newMIME("image/jpeg", ".jpg")
	newMIME("image/jxl", ".jxl")
	newMIME("image/jp2", ".jp2")
	newMIME("image/jpx", ".jpf")
	newMIME("image/jpm", ".jpm")
	newMIME("image/x-xpixmap", ".xpm")
	newMIME("image/bpg", ".bpg")
	newMIME("image/gif", ".gif")
	newMIME("image/webp", ".webp")
	newMIME("image/tiff", ".tiff")
	newMIME("image/bmp", ".bmp")
	newMIME("image/x-icon", ".ico")
	newMIME("image/x-icns", ".icns")
	newMIME("image/vnd.adobe.photoshop", ".psd")
	newMIME("image/heic", ".heic")
	newMIME("image/heic-sequence", ".heic")
	newMIME("image/heif", ".heif")
	newMIME("image/heif-sequence", ".heif")
	newMIME("image/vnd.radiance", ".hdr")
	newMIME("image/avif", ".avif")
	newMIME("audio/mpeg", ".mp3")
	newMIME("audio/flac", ".flac")
	newMIME("audio/midi", ".midi")
	newMIME("audio/ape", ".ape")
	newMIME("audio/musepack", ".mpc")
	newMIME("audio/wav", ".wav")
	newMIME("audio/aiff", ".aiff")
	newMIME("audio/basic", ".au")
	newMIME("audio/amr", ".amr")
	newMIME("audio/aac", ".aac")
	newMIME("audio/x-unknown", ".voc")
	newMIME("audio/mp4", ".mp4")
	newMIME("audio/x-m4a", ".m4a")
	newMIME("application/vnd.apple.mpegurl", ".m3u")
	newMIME("video/x-m4v", ".m4v")
	newMIME("video/mp4", ".mp4")
	newMIME("video/webm", ".webm")
	newMIME("video/mpeg", ".mpeg")
	newMIME("video/quicktime", ".mov")
	newMIME("video/quicktime", ".mqv")
	newMIME("video/3gpp", ".3gp")
	newMIME("video/3gpp2", ".3g2")
	newMIME("video/x-msvideo", ".avi")
	newMIME("video/x-flv", ".flv")
	newMIME("video/x-matroska", ".mkv")
	newMIME("video/x-ms-asf", ".asf")
	newMIME("application/vnd.rn-realmedia-vbr", ".rmvb")
	newMIME("application/x-java-applet", ".class")
	newMIME("application/x-shockwave-flash", ".swf")
	newMIME("application/x-chrome-extension", ".crx")
	newMIME("font/ttf", ".ttf")
	newMIME("font/woff", ".woff")
	newMIME("font/woff2", ".woff2")
	newMIME("font/otf", ".otf")
	newMIME("application/vnd.ms-fontobject", ".eot")
	newMIME("application/wasm", ".wasm")
	newMIME("application/octet-stream", ".shp")
	newMIME("application/octet-stream", ".shx")
	newMIME("application/x-dbf", ".dbf")
	newMIME("application/vnd.microsoft.portable-executable", ".exe")
	newMIME("application/x-sharedlib", ".so")
	newMIME("application/x-archive", ".a")
	newMIME("application/vnd.debian.binary-package", ".deb")
	newMIME("application/x-rpm", ".rpm")
	newMIME("application/dicom", ".dcm")
	newMIME("application/vnd.oasis.opendocument.text", ".odt")
	newMIME("application/vnd.oasis.opendocument.text-template", ".ott")
	newMIME("application/vnd.oasis.opendocument.spreadsheet", ".ods")
	newMIME("application/vnd.oasis.opendocument.spreadsheet-template", ".ots")
	newMIME("application/vnd.oasis.opendocument.presentation", ".odp")
	newMIME("application/vnd.oasis.opendocument.presentation-template", ".otp")
	newMIME("application/vnd.oasis.opendocument.graphics", ".odg")
	newMIME("application/vnd.oasis.opendocument.graphics-template", ".otg")
	newMIME("application/vnd.oasis.opendocument.formula", ".odf")
	newMIME("application/vnd.oasis.opendocument.chart", ".odc")
	newMIME("application/vnd.sun.xml.calc", ".sxc")
	newMIME("application/x-rar-compressed", ".rar")
	newMIME("image/vnd.djvu", ".djvu")
	newMIME("application/x-mobipocket-ebook", ".mobi")
	newMIME("application/x-ms-reader", ".lit")
	newMIME("application/vnd.sqlite3", ".sqlite")
	newMIME("image/vnd.dwg", ".dwg")
	newMIME("application/warc", ".warc")
	newMIME("application/vnd.nintendo.snes.rom", ".nes")
	newMIME("application/x-ms-shortcut", ".lnk")
	newMIME("application/x-mach-binary", ".macho")
	newMIME("audio/qcelp", ".qcp")
	newMIME("application/marc", ".mrc")
	newMIME("application/x-msaccess", ".mdb")
	newMIME("application/x-msaccess", ".accdb")
	newMIME("application/zstd", ".zst")
	newMIME("application/vnd.ms-cab-compressed", ".cab")
	newMIME("application/lzip", ".lz")
	newMIME("application/x-bittorrent", ".torrent")
	newMIME("application/x-cpio", ".cpio")
	newMIME("application/pkcs7-signature", ".p7s")
	newMIME("image/x-xcf", ".xcf")
	newMIME("image/x-gimp-pat", ".pat")
	newMIME("image/x-gimp-gbr", ".gbr")
	newMIME("application/vnd.adobe.xfdf", ".xfdf")
	newMIME("model/gltf-binary", ".glb")
}
