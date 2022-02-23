package repo

import (
	"github.com/ashadi-cc/resemble/v2/option"
	"github.com/ashadi-cc/resemble/v2/request"
	"github.com/ashadi-cc/resemble/v2/response"
)

// Clip represent clip method interface collections
type Clip interface {
	// All returns all clips by given projectUuid
	All(projectUuid string, page int, pageSize ...int) (response.Clips, error)

	// CreateSync create clip with synch method
	CreateSync(projectUuid string, data request.Payload) (response.Clip, error)

	// CreateAsync create clip with async method
	CreateAsync(projectUuid string, callbackUrl string, data request.Payload) (response.Clip, error)

	// Get returns clip by given projectuuid, clipuuid
	Get(projectUuid, uuid string) (response.Clip, error)

	// UpdateAsync update clip with asynch method
	UpdateAsync(projectUuid, uuid, callbackUrl string, data request.Payload) (response.Clip, error)

	// Delete delete clip by given projectuuid, clipuuid
	Delete(projectUuid, uuid string) (response.Message, error)

	// Stream returns stream a clip by given payload request
	Stream(data request.Payload, options ...option.ClipStream) (chan response.Metadata, chan []byte, chan bool, chan error)
}
