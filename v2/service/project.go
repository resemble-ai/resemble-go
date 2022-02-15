package service

import (
	"resemble/v2/repo"
	"resemble/v2/request"
	"resemble/v2/response"
)

// NewProject returns a new instance of repo.Project
func NewProject() repo.Project {
	return &project{}
}

type project struct{}

// All implements repo.Project.All method
func (p project) All(page, pageSize int) (response.Projects, error) {
	return response.Projects{}, nil
}

// Create implements repo.Project.Create method
func (p project) Create(data request.Payload) (response.Project, error) {
	return response.Project{}, nil
}

// Get implements repo.Project.Get method
func (p project) Get(uuid string) (response.Project, error) {
	return response.Project{}, nil
}

// Update implements repo.Project.Update method
func (p project) Update(uuid string, data request.Payload) (response.Project, error) {
	return response.Project{}, nil
}

// Delete implements repo.Project.Delete method
func (p project) Delete(uuid string) (response.Message, error) {
	return response.Message{}, nil
}
