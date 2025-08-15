package web

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed dist/*
var embedFS embed.FS

func GetFS() http.FileSystem {
	subFS, err := fs.Sub(embedFS, "dist")
	if err != nil {
		panic(err)
	}

	return http.FS(subFS)
}
