package repo

import (
	"resemble/v2/request"
	"resemble/v2/response"
)

// Clip represent clip method interface collections
type Clip interface {
	// Stream returns stream a clip by given syncServerUrl and payload
	Stream(syncServerUrl string, data request.Payload) (chan response.ClipStream, error)
}
