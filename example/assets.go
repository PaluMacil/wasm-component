package example

import "embed"

//go:embed web
var Assets embed.FS

// TODO: remove, currently used to list some files, but not needed
// Deprecated
//go:embed web/*.gohtml web/*/*.gohtml
var Templates embed.FS
