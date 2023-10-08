package imagedownloader

import (
	"net/http"
)

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	CommonImageContentTypeExtensions = map[string]string{
		/* ".jpeg":               true,
		".jpg":                true,
		".png":                true,
		".gif":                true,
		".bmp":                true,
		".x-ms-bmp":           true,
		".webp":               true,
		".svg+xml":            true,
		".x-icon":             true,
		".vnd.microsoft.icon": true,
		".tiff":               true,
		".vnd.radiance":       true,
		".jp2":                true, */
		"image/jpeg":               ".jpg",
		"image/jpg":                ".jpg",
		"image/png":                ".png",
		"image/gif":                ".gif",
		"image/bmp":                ".bmp",
		"image/x-ms-bmp":           ".bmp",
		"image/webp":               ".webp",
		"image/svg+xml":            ".svg",
		"image/x-icon":             ".ico",
		"image/vnd.microsoft.icon": ".ico",
		"image/tiff":               ".tiff",
		"image/vnd.radiance":       ".hdr",
		"image/jp2":                ".jp2",
	}
)
