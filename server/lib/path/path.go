package path

import (
	"path/filepath"
	"runtime"
)

var ()

func Root() string {
	_, b, _, _ := runtime.Caller(0)
	root := filepath.Join(filepath.Dir(b), "../..")
	return root
}

func WithRoot(s string) string {
	_, b, _, _ := runtime.Caller(0)
	root := filepath.Join(filepath.Dir(b), "../..")

	if s[0:1] == "/" || s[0:1] == "\\" {
		s = s[1:]
	}
	return root + "/" + s
}
