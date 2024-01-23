package utils

import (
	"log"
	"os"
	"strings"
	"time"
)

func Check(err error) {
	if err != nil {
		log.New(os.Stderr, "ERROR ", log.LstdFlags|log.Lshortfile).Panic(err)
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
)
