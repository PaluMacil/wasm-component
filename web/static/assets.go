package static

import "embed"

//go:embed *.js *.css *.map
var Templates embed.FS
