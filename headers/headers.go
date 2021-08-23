package headers

// Headers - contains the headers information
type Headers struct {
	headers map[string]string `json:"-"`
}

// New creates a new Headers instance.
// Core headers are added by default.
func New() *Headers {
	h := &Headers{headers: make(map[string]string)}
	h.SetCors()
	return h
}

// Add - adds new header to map
func (h *Headers) Add(key, values string) *Headers {
	h.headers[key] = values
	return h
}

// Set - overwrites all headers with new ones
func (h *Headers) Set(headers map[string]string) *Headers {
	h.headers = headers
	return h
}

// GetValueByKey - gets value form the headers by a specific key
func (h *Headers) GetValueByKey(key string) string {
	return h.headers[key]
}

// SetCors - adds cors headers
func (h *Headers) SetCors() *Headers {
	h.headers["Access-Control-Allow-Headers"] = "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token"
	h.headers["Access-Control-Allow-Methods"] = "GET,OPTIONS,POST,PUT,PATCH"
	h.headers["Access-Control-Allow-Origin"] = "*"
	return h
}

// Get - returns headers map
func (h *Headers) Get() map[string]string {
	return h.headers
}
