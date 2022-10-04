package datetime

import (
	"strings"
	"time"
)

var (
	DEFAULT = Format("default")
	HYPHEN  = Format("hyphen")
	SLASH   = Format("slash")
)

type Format string

func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func FormatDate(d string, f Format) string {
	date := strings.Replace(strings.Replace(d, "-", "", -1), "/", "", -1)

	if len(date) != 8 {
		return ""
	}

	if date[:1] != "2" {
		return ""
	}

	switch f {
	case HYPHEN:
		return date[0:4] + "-" + date[4:6] + "-" + date[6:8]
	case SLASH:
		return date[0:4] + "/" + date[4:6] + "/" + date[6:8]
	default:
		return date
	}
}

func Today(f Format) string {
	switch f {
	case HYPHEN:
		return time.Now().Format("2006-01-02")
	case SLASH:
		return time.Now().Format("2006/01/02")
	default:
		return time.Now().Format("20060102")
	}
}

func Now(f Format) string {
	switch f {
	case HYPHEN:
		return time.Now().Format("2006-01-02 15:04:05")
	case SLASH:
		return time.Now().Format("2006/01/02 15:04:05")
	default:
		return time.Now().Format("20060102150405")
	}
}
