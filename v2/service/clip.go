package service

import (
	"resemble/v2/repo"
	"resemble/v2/request"
	"resemble/v2/response"
)

// NewClip returns a new instance of repo.Client
func NewClip() repo.Clip {
	return &clip{}
}

type clip struct{}

// Stream implements repo.Clip.Stream method
func (c clip) Stream(syncServerUrl string, data request.Payload) (chan response.ClipStream, error) {
	return nil, nil
}
