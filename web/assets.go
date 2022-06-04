package web

import "embed"

//go:embed  *.js *.css

//go:embed *.gohtml
var Templates embed.FS
