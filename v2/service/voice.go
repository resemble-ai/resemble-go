package service

import "github.com/ashadi-cc/resemble/v2/repo"

// NewVoice returns a new instance of repo.Voice
func NewVoice() repo.Voice {
	return &voice{}
}

type voice struct{}
