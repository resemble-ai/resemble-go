package resemble

import (
	"github.com/ashadi-cc/resemble/v2/api"
	"github.com/ashadi-cc/resemble/v2/repo"
	"github.com/ashadi-cc/resemble/v2/service"
)

// NewClient returns a new instance resemble client
func NewClient(token string) *Client {
	apiClient := api.NewClient(token)
	return &Client{
		Project: service.NewProject(apiClient),
		Voice:   service.NewVoice(apiClient),
		Clip:    service.NewClip(apiClient),
	}
}

// Client represent resemble client
type Client struct {
	Project repo.Project
	Voice   repo.Voice
	Clip    repo.Clip
}
