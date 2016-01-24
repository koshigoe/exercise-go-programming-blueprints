package main

import (
	"time"
)

// message は1つのメッセージを表します
type message struct {
	Name string
	Message string
	When time.Time
	AvatarURL string
}
