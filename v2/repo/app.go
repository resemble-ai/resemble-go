package repo

// App base App interface methods
type App interface {
	// GetSyncServerUrl returns sync server url
	GetSyncServerUrl() string
}
