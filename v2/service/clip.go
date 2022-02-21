package service

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/ashadi-cc/resemble/v2/api"
	"github.com/ashadi-cc/resemble/v2/option"
	"github.com/ashadi-cc/resemble/v2/repo"
	"github.com/ashadi-cc/resemble/v2/request"
	"github.com/ashadi-cc/resemble/v2/response"
	"github.com/ashadi-cc/resemble/v2/util"

	"github.com/pkg/errors"
)

// NewClip returns a new instance of repo.Client
func NewClip(app repo.App, clientApi api.Operation) repo.Clip {
	return &clip{
		app:       app,
		clientApi: clientApi,
	}
}

type clip struct {
	clientApi api.Operation
	app       repo.App
}

//// All implements repo.Clip.All method
func (c clip) All(projectUuid string, page int, pageSize ...int) (response.Clips, error) {
	var clips response.Clips

	q := map[string]interface{}{}
	q["page"] = page
	if len(pageSize) > 0 {
		q["page_size"] = pageSize[0]
	}

	path := fmt.Sprintf("projects/%s/clips", projectUuid)
	resp, err := c.clientApi.Get(context.Background(), path, q)
	if err != nil {
		return clips, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return clips, errors.Wrap(err, "unable to read body")
	}

	if resp.StatusCode != http.StatusOK {
		return clips, fmt.Errorf("%s", string(body))
	}

	if err := json.Unmarshal(body, &clips); err != nil {
		return clips, err
	}

	return clips, nil
}

// CreateSync implements repo.Clip.CreateSync method
func (c clip) CreateSync(projectUuid string, data request.Payload) (response.Clip, error) {
	delete(data, "callback_uri")
	return c.create(projectUuid, data)
}

// CreateAsync implements repo.Clip.CreateAsync method
func (c clip) CreateAsync(projectUuid string, callbackUrl string, data request.Payload) (response.Clip, error) {
	data["callback_uri"] = callbackUrl
	return c.create(projectUuid, data)
}

// Get implements repo.Clip.Get method
func (c clip) Get(projectUuid, uuid string) (response.Clip, error) {
	path := fmt.Sprintf("projects/%s/clips/%s", projectUuid, uuid)

	var clip response.Clip
	resp, err := c.clientApi.Get(context.Background(), path)
	if err != nil {
		return clip, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return clip, errors.Wrap(err, "unable to read body")
	}

	if resp.StatusCode != http.StatusOK {
		return clip, fmt.Errorf("%s", string(body))
	}

	if err := json.Unmarshal(body, &clip); err != nil {
		return clip, err
	}

	return clip, nil
}

// UpdateAsync implements repo.Clip.UpdateAsync method
func (c clip) UpdateAsync(projectUuid, uuid, callbackUrl string, data request.Payload) (response.Clip, error) {
	path := fmt.Sprintf("projects/%s/clips/%s", projectUuid, uuid)
	data["callback_uri"] = callbackUrl

	var clip response.Clip
	resp, err := c.clientApi.Put(context.Background(), path, data)
	if err != nil {
		return clip, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return clip, errors.Wrap(err, "unable to read body")
	}

	if resp.StatusCode != http.StatusOK {
		return clip, fmt.Errorf("%s", string(body))
	}

	if err := json.Unmarshal(body, &clip); err != nil {
		return clip, err
	}

	return clip, nil
}

// Delete implements repo.Clip.Delete method
func (c clip) Delete(projectUuid, uuid string) (response.Message, error) {
	path := fmt.Sprintf("projects/%s/clips/%s", projectUuid, uuid)
	var message response.Message
	resp, err := c.clientApi.Delete(context.Background(), path, nil)
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

// Stream implements repo.Clip.Stream method
func (c clip) Stream(data request.Payload, options ...option.ClipStream) (chan response.ClipStream, error) {
	resp, err := c.clientApi.Stream(context.Background(), c.app.GetSyncServerUrl(), data)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "unable to read body")
		}
		return nil, fmt.Errorf("%s", string(body))
	}

	opt := option.ClipStream{}
	if options != nil {
		opt = options[0]
	}
	opt.Parse()

	decoder, err := util.NewStreamDecoder(opt.BufferSize, !opt.WithWavHeader)
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReaderSize(resp.Body, opt.ChunkSize)
	clipStream := make(chan response.ClipStream)

	go c.decodeChunk(reader, decoder, clipStream)

	return clipStream, nil
}

func (c clip) decodeChunk(reader *bufio.Reader, decoder *util.StreamDecoder, clipStream chan response.ClipStream) {
	defer close(clipStream)

	for {
		chunk, err := reader.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				clipStream <- response.ClipStream{Chunk: nil, Err: err}
			}
			break
		}

		decoder.DecodeChunk(chunk)
		if buffer := decoder.FlushBuffer(false); buffer != nil {
			clipStream <- response.ClipStream{Chunk: buffer}
		}

	}

	for {
		if buffer := decoder.FlushBuffer(false); buffer != nil {
			clipStream <- response.ClipStream{Chunk: buffer}
		} else {
			return
		}
	}
}

func (c clip) create(projectUuid string, data request.Payload) (response.Clip, error) {
	path := fmt.Sprintf("projects/%s/clips", projectUuid)

	var clip response.Clip
	resp, err := c.clientApi.Post(context.Background(), path, data)
	if err != nil {
		return clip, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return clip, errors.Wrap(err, "unable to read body")
	}

	if resp.StatusCode != http.StatusOK {
		return clip, fmt.Errorf("%s", string(body))
	}

	if err := json.Unmarshal(body, &clip); err != nil {
		return clip, err
	}

	return clip, nil
}
