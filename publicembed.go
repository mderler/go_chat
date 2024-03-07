package gochat

import (
	"embed"
)

//go:embed public/*
var Public embed.FS

//go:embed migrations/*.sql
var EmbedMigrations embed.FS
