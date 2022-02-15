package resemble

import (
	"resemble/v2/api"
	"resemble/v2/repo"
	"resemble/v2/service"
)

// NewClient returns a new instance resemble client
func NewClient(token string) *Client {
	return &Client{
		Project: service.NewProject(),
		Voice:   service.NewVoice(),
		Clip:    service.NewClip(api.NewClient(token)),
	}
}

// Client represent resemble client
type Client struct {
	Project repo.Project
	Voice   repo.Voice
	Clip    repo.Clip
}
