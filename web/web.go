package web

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed dist
var eFS embed.FS

func GetFS() http.FileSystem {
	subFS, err := fs.Sub(eFS, "dist")
	if err != nil {
		panic(err)
	}

	return http.FS(subFS)
}
