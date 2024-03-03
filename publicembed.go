package gochat

import _ "embed"

//go:embed public/style.css
var Styles string

//go:embed public/htmx.min.js
var HTMX string
