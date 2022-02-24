package repo

import (
	"github.com/resemble-ai/resemble-go/v2/request"
	"github.com/resemble-ai/resemble-go/v2/response"
)

// Voice represent voice method interface collections
type Voice interface {
	// All returns all voices by given page and pagesize
	All(page int, pageSize ...int) (response.Voices, error)

	// Create create new voice by given data payload
	Create(data request.Payload) (response.Voice, error)

	// Get returns voice by given uuid
	Get(uuid string) (response.Voice, error)

	// Update update voice by given uuid
	Update(uuid string, data request.Payload) (response.Voice, error)

	// Delete delete voice by given uuid
	Delete(uuid string) (response.Message, error)

	// Build build voice by given uuid
	Build(uuid string) (response.Message, error)
}
