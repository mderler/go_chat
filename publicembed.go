package gochat

import (
	"embed"
	_ "embed"
)

//go:embed public/style.css
var Styles string

//go:embed public/htmx.min.js
var HTMX string

//go:embed migrations/*.sql
var EmbedMigrations embed.FS
