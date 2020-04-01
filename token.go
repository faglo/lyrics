package main

import (
	"io/ioutil"
	"os"
	"runtime"
)

func getPath() (path string, dirs string) {
	switch runtime.GOOS {
	case "windows":
		dirs = os.Getenv("APPDATA") + "/lyrics/"
	case "linux":
	case "freebsd":
	case "openbsd":
	case "darwin":
		dirs = os.Getenv("HOME") + "/.config/lyrics/"
	default:
		panic("unsupported system")
	}
	path = dirs + "genius"
	return
}

func token() (token string, ok bool) {
	if t := os.Getenv("GENIUS"); t != "" {
		return t, true
	}
	path, _ := getPath()
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	token, ok = string(b), true
	return
}

func setToken(token string) {
	path, dirs := getPath()
	err := os.MkdirAll(dirs, os.ModePerm)
	checkErr(err)
	err = ioutil.WriteFile(path, []byte(token), 0644)
	checkErr(err)
}

func removeToken() {
	path, _ := getPath()
	err := os.Remove(path)
	checkErr(err)
}
