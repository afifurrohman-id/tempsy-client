package utils

import (
	"log"
	"strings"
	"time"
)

func Check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// FormatDate date parameter is time in milliseconds
func FormatDate(date int64) string {
	dateWithDot := strings.Split(time.UnixMilli(date).Local().String(), "+")[0]
	return strings.Split(dateWithDot, ".")[0]
}

const (
	HeaderAutoDeleteAt      = "File-Auto-Delete-At"
	HeaderPrivateUrlExpires = "File-Private-Url-Expires"
	HeaderIsPublic          = "File-Is-Public"
	HeaderFileName          = "File-Name"
	HeaderRealIp            = "Real-IP"
	BearerPrefix            = "Bearer "
	DefaultApiTimeout       = 20 * time.Second
)
