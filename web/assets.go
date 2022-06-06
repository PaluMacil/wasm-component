package web

import "embed"

//go:embed static
var Assets embed.FS

//go:embed *.gohtml */*.gohtml
var Templates embed.FS
