package shared

import (
	"runtime"
	"strings"
)

func FormatPath(path string) string {
	if runtime.GOOS == "windows" {
		return strings.ReplaceAll(path, "/", "\\")
	} else {
		return strings.ReplaceAll(path, "\\", "/")
	}
}
