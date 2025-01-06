package internal

import (
	"embed"
)

// Templates is the embedded filesystem containing the templates.
//
//go:embed templates/*
var Templates embed.FS
