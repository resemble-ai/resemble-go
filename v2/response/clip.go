package response

// ClipStream represent clip stream response data model
type ClipStream struct {
	Chunk []byte
	Err   error
}
