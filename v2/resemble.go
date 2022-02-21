package resemble

import (
	"github.com/ashadi-cc/resemble/v2/api"
	"github.com/ashadi-cc/resemble/v2/repo"
	"github.com/ashadi-cc/resemble/v2/service"
)

// NewClient returns a new instance resemble client
func NewClient(token string) *Client {
	apiClient := api.NewClient(token)
	client := &Client{
		Project:   service.NewProject(apiClient),
		Voice:     service.NewVoice(apiClient),
		Recording: service.NewRecording(apiClient),
	}

	client.Clip = service.NewClip(client, apiClient)
	return client
}

// Client represent resemble client
type Client struct {
	Project       repo.Project
	Voice         repo.Voice
	Clip          repo.Clip
	Recording     repo.Recording
	syncServerUrl string
}

// SetSyncServerUrl set sync server url
func (c *Client) SetSyncServerUrl(uri string) {
	c.syncServerUrl = uri
}

// GetSyncServerUrl returns sync server url value
func (c *Client) GetSyncServerUrl() string {
	return c.syncServerUrl
}
