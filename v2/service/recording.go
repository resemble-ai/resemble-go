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
	"github.com/pkg/errors"
)

// NewRecording returns a new instance of repo.Recording
func NewRecording(clientApi api.Operation) repo.Recording {
	return &recording{
		clientApi: clientApi,
	}
}

type recording struct {
	clientApi api.Operation
}

// All implements repo.Recording.All method
func (r recording) All(uuid string, page int, pageSize ...int) (response.Recordings, error) {
	q := map[string]interface{}{}
	q["page"] = page
	if len(pageSize) > 0 {
		q["page_size"] = pageSize[0]
	}

	path := fmt.Sprintf("voices/%s/recordings", uuid)
	var recordings response.Recordings
	resp, err := r.clientApi.Get(context.Background(), path, q)
	if err != nil {
		return recordings, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return recordings, errors.Wrap(err, "unable to read body")
	}

	if resp.StatusCode != http.StatusOK {
		return recordings, fmt.Errorf("%s", string(body))
	}

	if err := json.Unmarshal(body, &recordings); err != nil {
		return recordings, err
	}

	return recordings, nil
}

// Create implements repo.Recording.Create method
func (r recording) Create(voiceuuid string, filePath string, data request.Payload) (response.Recording, error) {
	data["file"] = fmt.Sprintf("@%s", filePath)
	path := fmt.Sprintf("voices/%s/recordings", voiceuuid)

	payload := map[string]string{}
	for k, v := range data {
		payload[k] = fmt.Sprint(v)
	}

	var recording response.Recording
	resp, err := r.clientApi.PostForm(context.Background(), path, payload)
	if err != nil {
		return recording, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return recording, errors.Wrap(err, "unable to read body")
	}

	if resp.StatusCode != http.StatusOK {
		return recording, fmt.Errorf("%s", string(body))
	}

	if err := json.Unmarshal(body, &recording); err != nil {
		return recording, err
	}

	return recording, nil
}

// Get implements repo.Recording.Get method
func (r recording) Get(voiceuuid, uuid string) (response.Recording, error) {
	path := fmt.Sprintf("voices/%s/recordings/%s", voiceuuid, uuid)
	var recording response.Recording
	resp, err := r.clientApi.Get(context.Background(), path)
	if err != nil {
		return recording, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return recording, errors.Wrap(err, "unable to read body")
	}

	if resp.StatusCode != http.StatusOK {
		return recording, fmt.Errorf("%s", string(body))
	}

	if err := json.Unmarshal(body, &recording); err != nil {
		return recording, err
	}

	return recording, nil
}

// Update implements repo.Recording.Update method
func (r recording) Update(voiceuuid, uuid string, data request.Payload) (response.Recording, error) {
	path := fmt.Sprintf("voices/%s/recordings/%s", voiceuuid, uuid)
	var recording response.Recording
	resp, err := r.clientApi.Put(context.Background(), path, data)
	if err != nil {
		return recording, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return recording, errors.Wrap(err, "unable to read body")
	}

	if resp.StatusCode != http.StatusOK {
		return recording, fmt.Errorf("%s", string(body))
	}

	if err := json.Unmarshal(body, &recording); err != nil {
		return recording, err
	}

	return recording, nil
}

// Delete implements repo.Recording.Delete method
func (r recording) Delete(voiceuuid, uuid string) (response.Message, error) {
	path := fmt.Sprintf("voices/%s/recordings/%s", voiceuuid, uuid)

	var message response.Message
	resp, err := r.clientApi.Delete(context.Background(), path, nil)
	if err != nil {
		return message, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return message, errors.Wrap(err, "unable to read body")
	}

	if resp.StatusCode != http.StatusOK {
		return message, fmt.Errorf("%s", string(body))
	}

	if err := json.Unmarshal(body, &message); err != nil {
		return message, err
	}

	return message, nil
}
