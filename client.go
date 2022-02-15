package resemble

import (
	"resemble/repo"
	"resemble/service"
)

// NewClient returns a new instance resemble client
func NewClient(version, token string) *Client {
	return &Client{
		Project: service.NewProject(),
		Voice:   service.NewVoice(),
		Clip:    service.NewClip(),
	}
}

// Client represent resemble client
type Client struct {
	Project repo.Project
	Voice   repo.Voice
	Clip    repo.Clip
}
