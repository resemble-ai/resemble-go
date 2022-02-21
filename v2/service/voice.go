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

// NewVoice returns a new instance of repo.Voice
func NewVoice(apiClient api.Operation) repo.Voice {
	return &voice{
		clientApi: apiClient,
	}
}

type voice struct {
	clientApi api.Operation
}

// All implements repo.Voice.All method
func (v voice) All(page int, pageSize ...int) (response.Voices, error) {
	q := map[string]interface{}{}
	q["page"] = page
	if len(pageSize) > 0 {
		q["page_size"] = pageSize[0]
	}

	var voices response.Voices
	resp, err := v.clientApi.Get(context.Background(), "voices", q)
	if err != nil {
		return voices, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return voices, errors.Wrap(err, "unable to read body")
	}

	if resp.StatusCode != http.StatusOK {
		return voices, util.NewApiError(body, "voices", resp.StatusCode, resp.Request.Method)
	}

	if err := json.Unmarshal(body, &voices); err != nil {
		return voices, err
	}

	return voices, nil
}

// Create implements repo.Voice.Create method
func (v voice) Create(data request.Payload) (response.Voice, error) {
	var voice response.Voice
	resp, err := v.clientApi.Post(context.Background(), "voices", data)
	if err != nil {
		return voice, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return voice, errors.Wrap(err, "unable to read body")
	}

	if resp.StatusCode != http.StatusOK {
		return voice, util.NewApiError(body, "voices", resp.StatusCode, resp.Request.Method)
	}

	if err := json.Unmarshal(body, &voice); err != nil {
		return voice, err
	}

	return voice, nil
}

// Get implements repo.Voice.Get method
func (v voice) Get(uuid string) (response.Voice, error) {
	path := fmt.Sprintf("voices/%s", uuid)
	var voice response.Voice
	resp, err := v.clientApi.Get(context.Background(), path)
	if err != nil {
		return voice, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return voice, errors.Wrap(err, "unable to read body")
	}

	if resp.StatusCode != http.StatusOK {
		return voice, util.NewApiError(body, path, resp.StatusCode, resp.Request.Method)
	}

	if err := json.Unmarshal(body, &voice); err != nil {
		return voice, err
	}

	return voice, nil
}

// Update implements repo.Voice.Update method
func (v voice) Update(uuid string, data request.Payload) (response.Voice, error) {
	path := fmt.Sprintf("voices/%s", uuid)
	var voice response.Voice
	resp, err := v.clientApi.Put(context.Background(), path, data)
	if err != nil {
		return voice, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return voice, errors.Wrap(err, "unable to read body")
	}

	if resp.StatusCode != http.StatusOK {
		return voice, util.NewApiError(body, path, resp.StatusCode, resp.Request.Method)
	}

	if err := json.Unmarshal(body, &voice); err != nil {
		return voice, err
	}

	return voice, nil
}

// Delete implements repo.Voice.Delete method
func (v voice) Delete(uuid string) (response.Message, error) {
	path := fmt.Sprintf("voices/%s", uuid)
	var message response.Message
	resp, err := v.clientApi.Delete(context.Background(), path, nil)
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

// Build implements repo.Voice.Build method
func (v voice) Build(uuid string) (response.Message, error) {
	path := fmt.Sprintf("voices/%s/build", uuid)
	var message response.Message
	resp, err := v.clientApi.Post(context.Background(), path, nil)
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
