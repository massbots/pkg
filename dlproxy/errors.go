package dlproxy

import (
	"regexp"
	"strings"
)

var (
	ErrDataBlocks          = "Did not get any data blocks"
	ErrGivingUp            = "giving up after 0 fragment retries"
	ErrNotAvailable        = "This video is not available"
	ErrNotAvailableCountry = "The uploader has not made this video available in your country"
	ErrContainsContent     = "This video contains content from"
	ErrRemoved             = "This video has been removed by the user"
	ErrViolating           = "This video has been removed for violating"
	ErrPrivate             = "This video is private"
	ErrUserIsBlocked       = "[400 USER_IS_BLOCKED]" // prefix
)

var errorKeys = map[string]string{
	ErrDataBlocks:          "cant_download",
	ErrGivingUp:            "cant_download",
	ErrNotAvailable:        "not_available",
	ErrNotAvailableCountry: "country",
	ErrContainsContent:     "content",
	ErrRemoved:             "removed",
	ErrViolating:           "violating",
	ErrPrivate:             "private",
}

var reContentOwner = regexp.MustCompile(ErrContainsContent + ` (.*), who`)

type Error string

func (err Error) Key() string {
	for k, v := range errorKeys {
		if strings.Contains(string(err), k) {
			return v
		}
	}
	return "not_available"
}

func (err Error) ContentOwner() string {
	if err.Key() == "content" {
		s := reContentOwner.FindStringSubmatch(string(err))
		if len(s) > 1 {
			return s[1]
		}
	}
	return ""
}

func (err Error) Error() string {
	return string(err)
}

func IsUserBlocked(err Error) bool {
	return strings.HasPrefix(string(err), ErrUserIsBlocked)
}
