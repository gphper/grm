package web

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
)

//go:embed dist
var StaticPath embed.FS

var StaticsFs http.FileSystem

func init() {
	static, err := fs.Sub(StaticPath, "dist")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(static)

	StaticsFs = http.FS(static)
}
