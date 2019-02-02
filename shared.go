package example

// IndexRequest container.
// The URL is the url to index content from.
type IndexRequest struct {
	URL string `json:"url"`
}

const (
	// IndexRequestJob queue name52415697
	IndexRequestJob = "IndexRequests"
)
