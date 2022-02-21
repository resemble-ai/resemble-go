package repo

import (
	"github.com/ashadi-cc/resemble/v2/request"
	"github.com/ashadi-cc/resemble/v2/response"
)

// Recording represent recording interface method collections
type Recording interface {
	// All returns all recording by given voice uuid
	All(voiceuuid string, page int, pageSize ...int) (response.Recordings, error)

	// Create create new recording by given data payload
	Create(voiceuuid string, filePath string, data request.Payload) (response.Recording, error)

	// Get returns recording by given voice uuid and recording uuid
	Get(voiceuuid, uuid string) (response.Recording, error)

	// Update update recording by given voice uuid and recoding uuid
	Update(voiceuuid, uuid string, data request.Payload) (response.Recording, error)

	// Delete delete recording by given voice uuid
	Delete(voiceuuid, uuid string) (response.Message, error)
}
