package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ashadi-cc/resemble/v2/api"
	"github.com/ashadi-cc/resemble/v2/repo"
	"github.com/ashadi-cc/resemble/v2/request"
	"github.com/ashadi-cc/resemble/v2/response"
	"github.com/ashadi-cc/resemble/v2/util"
	"github.com/pkg/errors"
)

// NewProject returns a new instance of repo.Project
func NewProject(clientApi api.Operation) repo.Project {
	return &project{
		clientApi: clientApi,
	}
}

type project struct {
	clientApi api.Operation
}

// All implements repo.Project.All method
func (p project) All(page int, pageSize ...int) (response.Projects, error) {
	q := map[string]interface{}{}
	q["page"] = page
	if len(pageSize) > 0 {
		q["page_size"] = pageSize[0]
	}

	var projects response.Projects
	resp, err := p.clientApi.Get(context.Background(), "projects", q)
	if err != nil {
		return projects, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return projects, errors.Wrap(err, "unable to read body")
	}

	if resp.StatusCode != http.StatusOK {
		return projects, util.NewApiError(body, "projects", resp.StatusCode, resp.Request.Method)
	}

	if err := json.Unmarshal(body, &projects); err != nil {
		return projects, err
	}

	return projects, nil
}

// Create implements repo.Project.Create method
func (p project) Create(data request.Payload) (response.Project, error) {
	var project response.Project
	resp, err := p.clientApi.Post(context.Background(), "projects", data)
	if err != nil {
		return project, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return project, errors.Wrap(err, "unable to read body")
	}

	if resp.StatusCode != http.StatusOK {
		return project, util.NewApiError(body, "projects", resp.StatusCode, resp.Request.Method)
	}

	if err := json.Unmarshal(body, &project); err != nil {
		return project, err
	}

	return project, nil
}

// Get implements repo.Project.Get method
func (p project) Get(uuid string) (response.Project, error) {
	path := fmt.Sprintf("projects/%s", uuid)
	var project response.Project
	resp, err := p.clientApi.Get(context.Background(), path)
	if err != nil {
		return project, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return project, errors.Wrap(err, "unable to read body")
	}

	if resp.StatusCode != http.StatusOK {
		return project, util.NewApiError(body, path, resp.StatusCode, resp.Request.Method)
	}

	if err := json.Unmarshal(body, &project); err != nil {
		return project, err
	}

	return project, nil
}

// Update implements repo.Project.Update method
func (p project) Update(uuid string, data request.Payload) (response.Project, error) {
	path := fmt.Sprintf("projects/%s", uuid)
	var project response.Project
	resp, err := p.clientApi.Put(context.Background(), path, data)
	if err != nil {
		return project, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return project, errors.Wrap(err, "unable to read body")
	}

	if resp.StatusCode != http.StatusOK {
		return project, util.NewApiError(body, path, resp.StatusCode, resp.Request.Method)
	}

	if err := json.Unmarshal(body, &project); err != nil {
		return project, err
	}

	return project, nil
}

// Delete implements repo.Project.Delete method
func (p project) Delete(uuid string) (response.Message, error) {
	path := fmt.Sprintf("projects/%s", uuid)
	var message response.Message
	resp, err := p.clientApi.Delete(context.Background(), path, nil)
	if err != nil {
		return message, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return message, errors.Wrap(err, "unable to read body")
	}

	if resp.StatusCode != http.StatusOK {
		return message, util.NewApiError(body, path, resp.StatusCode, resp.Request.Method)
	}

	if err := json.Unmarshal(body, &message); err != nil {
		return message, err
	}

	return message, nil
}
