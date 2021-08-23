package response

// Response
type Response interface {
	StatusCode() int
	GetBody() ([]byte, error)
	GetHeaders() map[string]string
	Error() string
	GetData() interface{}
}
