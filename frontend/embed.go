package frontend

import "embed"

//go:embed all:dist
var FrontendContentFS embed.FS
