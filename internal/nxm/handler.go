package nxm

import (
	log "github.com/sirupsen/logrus"
	"net/url"
	"strings"
)

func HandleNxmURL(parsedURL *url.URL) NXMInfo {
	switch parsedURL.Scheme {
	case "nxm":
		return handleNXM(parsedURL)
	default:
		log.WithField("schema", parsedURL.Scheme).Debug("Unknown URL scheme")
	}

	return NXMInfo{}
}

func handleNXM(u *url.URL) NXMInfo {
	log.WithField("url", u.String()).WithField("path", u.Path).WithField("query", u.RawQuery).Debug("Handling NMM URL")

	pathParts := strings.Split(u.Path, "/")

	return NXMInfo{
		ModID:   pathParts[2],
		FileID:  pathParts[4],
		Key:     u.Query().Get("key"),
		Expires: u.Query().Get("expires"),
	}
}
