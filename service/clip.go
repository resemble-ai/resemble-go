package service

import (
	"resemble/repo"
	"resemble/request"
	"resemble/response"
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
