package repo

import (
	"github.com/resemble-ai/resemble-go/v2/request"
	"github.com/resemble-ai/resemble-go/v2/response"
)

// Project represent project interface method collections
type Project interface {
	// All returns all project by given page and pagesize
	All(page int, pageSize ...int) (response.Projects, error)

	// Create create new project by given data payload
	Create(data request.Payload) (response.Project, error)

	// Get returns project by given uuid
	Get(uuid string) (response.Project, error)

	// Update update project by given uuid
	Update(uuid string, data request.Payload) (response.Project, error)

	// Delete delete project by given uuid
	Delete(uuid string) (response.Message, error)
}
