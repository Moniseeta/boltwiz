package ui

import "embed"

var (
	//go:embed dist/*
	WebContent embed.FS
)
