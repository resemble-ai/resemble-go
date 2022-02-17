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

	// Stream returns stream a clip by given syncServerUrl and payload
	Stream(syncServerUrl string, data request.Payload, options ...option.ClipStream) (chan response.ClipStream, error)
}
