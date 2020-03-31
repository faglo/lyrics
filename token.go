package main

import (
	"io/ioutil"
	"os"
	"runtime"
)

func token() (token string, ok bool) {
	if t := os.Getenv("GENIUS_TOKEN"); t != "" {
		return t, true
	}

	var path string

	switch runtime.GOOS {
	case "windows":
		path = os.Getenv("APPDATA") + "/lyrics/genius"
	case "linux":
	case "freebsd":
	case "openbsd":
	case "darwin":
		path = os.Getenv("HOME") + "/.config/lyrics/genius"
	default:
		panic("unsupported system")
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	token, ok = string(b), true
	return
}
