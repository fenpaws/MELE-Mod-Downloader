package nexusmods

import "errors"

var (
	ErrInvalidKeyOrExpireTime = errors.New("400: provided key and expire time isn't correct for this user/file")
	ErrPermissionDenied       = errors.New("403: you don't have permission to get download links from the API without visiting nexusmods.com - this is for premium users only")
	ErrFileNotFound           = errors.New("404: file not found")
	ErrLinkExpired            = errors.New("410: this link has expired - please visit the mod page again to get a new link")
)
