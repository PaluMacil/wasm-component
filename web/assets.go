package web

import "embed"

//go:embed *.gohtml
var Templates embed.FS
