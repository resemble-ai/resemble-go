package service

import "resemble/repo"

// NewVoice returns a new instance of repo.Voice
func NewVoice() repo.Voice {
	return &voice{}
}

type voice struct{}
